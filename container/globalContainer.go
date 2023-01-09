package container

import (
	"crypto/sha1"
	"fmt"
	"sync"

	"github.com/binarycurious/light-container/telemetry"
)

// GlobalContainer :  is an implementation of a Global - Container interface
type GlobalContainer struct {
	logger   telemetry.Logger
	state    *GlobalState
	routines map[RoutineKey]Routine
	inChans  map[string]chan RoutineMsg
	outChans map[string]<-chan RoutineMsg
	subChans map[string][]chan RoutineMsg

	running       bool
	containerLock sync.Mutex
}

// NewDefaultContainer Setup a default GlobalContainer for registering Routines
func (c GlobalContainer) NewDefaultContainer(state *GlobalState) Container {
	return c.NewContainer(state, nil)
}

// NewContainer - initialize a new IoC container, nil logger will create a new logger based on the globalstate settings
func (c GlobalContainer) NewContainer(state *GlobalState, logger telemetry.Logger) Container {

	c.containerLock.Lock()

	if logger == nil {
		logger = telemetry.Logger(&GlobalLogger{hardfail: state.settings.Hardfail})
	}
	c.logger = logger

	/*Setup channel register vars*/
	c.inChans = make(map[string]chan RoutineMsg, 100)
	c.subChans = make(map[string][]chan RoutineMsg, 100)

	validSate := true

	if !state.intitialized {
		msg := "FATAL: Attempt to initialize container state with an un-initialized global state (call NewState())"
		c.logger.LogFatal(&msg)
		validSate = false
	}

	if c.state != nil && c.state.intitialized {
		msg := "FATAL: Attempt to set container state after already intialized"
		c.logger.LogFatal(&msg)
		validSate = false
	}

	if validSate {
		c.state = state
	}

	c.containerLock.Unlock()

	return Container(&c)
}

// GetRoutineKey : Get a key for given routine name (This does not take into account name conflicts / duplicates)
func (c *GlobalContainer) GetRoutineKey(routineName *string) *RoutineKey {
	return &RoutineKey{
		name: routineName,
		key:  fmt.Sprintf("%x", sha1.Sum([]byte(*routineName)))}
}

/* Continer Impls */

// GetState @impl
func (c *GlobalContainer) GetState() *GlobalState {
	return c.state
}

// GetLogger @impl
func (c *GlobalContainer) GetLogger() telemetry.Logger {
	return c.logger
}

// AddRoutine : impl of Container.AddRoutineWithName (will modify routineName if there is a conflict)
func (c *GlobalContainer) AddRoutine(routine Routine) *RoutineKey {
	return c.AddRoutineWithName(routine.GetName(), routine)
}

// AddRoutineWithName : impl of Container.AddRoutineWithName (will modify routineName if there is a conflict)
func (c *GlobalContainer) AddRoutineWithName(routineName *string, routine Routine) *RoutineKey {
	c.containerLock.Lock()

	if len(c.routines) == 0 {
		c.routines = make(map[RoutineKey]Routine, 10)
	}
	rKey := c.GetRoutineKey(routineName)
	_, retry := c.routines[*rKey]

	retries := 0
	for retry {
		retries++
		*routineName = fmt.Sprintf("%s_%d", *routineName, retries)
		rKey = c.GetRoutineKey(routineName)
		_, retry = c.routines[*rKey]
	}

	c.routines[*rKey] = routine

	c.containerLock.Unlock()
	return rKey
}

// Execute : impl of container Execute function
func (c *GlobalContainer) Execute(key *RoutineKey) error {
	c.containerLock.Lock()

	ctx := RoutineContext{}.NewRoutineContext(key, Container(c))
	rcvr := make(chan RoutineMsg)
	c.inChans[key.key] = rcvr
	sc, err := c.routines[*key].Subscribe()
	if err != nil {
		return err
	}
	c.outChans[key.key] = sc

	c.containerLock.Unlock()

	go (c.routines[*key]).Execute(Context(ctx), rcvr)

	return nil
}

// Send @impl - send to a routine channel
func (c *GlobalContainer) Send(key *RoutineKey, msg RoutineMsg) error {

	if c.inChans[key.key] == nil {
		return fmt.Errorf("Cannot sending msg to nil channel for routineKey %#v", key)
	}

	c.inChans[key.key] <- msg
	return nil
}

// Subscribe @impl - subscribe to a given routine channel
func (c *GlobalContainer) Subscribe(key *RoutineKey) (<-chan RoutineMsg, error) {
	c.containerLock.Lock()
	var sc chan RoutineMsg
	if c.outChans[key.key] != nil {
		sc = make(chan RoutineMsg)
		c.subChans[key.key] = append(c.subChans[key.key], sc)
	} else {
		return nil, fmt.Errorf("No out channel matching key")
	}
	c.containerLock.Unlock()

	return sc, nil
}

// Start @impl Continer.Start - Spins up all container routines and msg wiring
func (c *GlobalContainer) Start() {

	go func() {
		for k := range c.routines {
			c.Execute(&k)
		}
	}()

	c.running = true
	for c.running == true {
		for k, oc := range c.outChans {
			select {
			case msg := <-oc:
				for i := range c.subChans[k] {
					c.subChans[k][i] <- msg
				}
			}
		}
	}
	for _, val := range c.inChans {
		close(val)
	}
	for _, arr := range c.subChans {
		for i := range arr {
			close(arr[i])
		}
	}
}

// Stop @impl Container.Stop - Stops the container routines
func (c *GlobalContainer) Stop() {
	c.running = false
}
