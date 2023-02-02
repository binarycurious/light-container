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
	keys     map[string]RoutineKey // map of name : routinekey
	routines map[string]Routine
	inChans  map[string]chan RoutineMsg
	outChans map[string]chan RoutineMsg
	subChans map[string][]chan RoutineMsg

	running       bool
	containerLock sync.Mutex
}

// NewDefaultContainer Setup a default GlobalContainer for registering Routines
func NewDefaultContainer(state *GlobalState) Container {
	return NewContainer(state, nil)
}

// NewContainer - initialize a new IoC container, nil logger will create a new logger based on the globalstate settings
func NewContainer(state *GlobalState, logger telemetry.Logger) Container {
	c := GlobalContainer{}

	c.containerLock.Lock()

	if logger == nil {
		logger = telemetry.Logger(&GlobalLogger{hardfail: state.settings.Hardfail, loglevel: state.settings.LogLevel})
	}
	c.logger = logger

	/*Setup channel register vars*/
	c.inChans = make(map[string]chan RoutineMsg, 100)
	c.outChans = make(map[string]chan RoutineMsg, 100)
	c.subChans = make(map[string][]chan RoutineMsg, 100)

	validSate := true

	if !state.intitialized {
		msg := "FATAL: Attempt to initialize container state with an un-initialized global state (call NewState())"
		c.logger.LogFatal(msg)
		validSate = false
	}

	if c.state != nil && c.state.intitialized {
		msg := "FATAL: Attempt to set container state after already intialized"
		c.logger.LogFatal(msg)
		validSate = false
	}

	if validSate {
		c.state = state
	}

	c.containerLock.Unlock()

	return Container(&c)
}

// GenRoutineKey : Get a key for given routine name (This does not take into account name conflicts / duplicates)
func (c *GlobalContainer) GenRoutineKey(routineName *string) *RoutineKey {
	return &RoutineKey{
		name: routineName,
		key:  fmt.Sprintf("%x", sha1.Sum([]byte(*routineName)))}
}

// GetRoutineKey -
func (c *GlobalContainer) GetRoutineKey(routineName *string) *RoutineKey {
	k := c.keys[*routineName]
	return &k
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
func (c *GlobalContainer) AddRoutine(routine Routine) RoutineKey {
	return c.AddNamedRoutine(routine.GetName(), routine)
}

// AddNamedRoutine : impl of Container.AddNamedRoutine (will modify routineName if there is a conflict)
func (c *GlobalContainer) AddNamedRoutine(routineName *string, routine Routine) RoutineKey {

	c.containerLock.Lock()

	if len(c.routines) == 0 {
		c.keys = make(map[string]RoutineKey, 10)
		c.routines = make(map[string]Routine, 10)
	}
	rKey := c.GenRoutineKey(routineName)
	_, retry := c.routines[rKey.key]

	retries := 0
	for retry {
		retries++
		*routineName = fmt.Sprintf("%s_%d", *routineName, retries)
		rKey = c.GetRoutineKey(routineName)
		_, retry = c.routines[rKey.key]
	}

	c.keys[*rKey.name] = *rKey
	c.routines[rKey.key] = routine

	rInCh := make(chan RoutineMsg)
	c.inChans[rKey.key] = rInCh

	rOutCh := make(chan RoutineMsg)
	c.outChans[rKey.key] = rOutCh

	c.containerLock.Unlock()
	return *rKey
}

// Execute : impl of container Execute function
func (c *GlobalContainer) Execute(key *RoutineKey, wg *sync.WaitGroup) error {

	rInCh := c.inChans[key.key]
	rOutCh := c.outChans[key.key]

	ctx := NewRoutineContext(key, Container(c), rInCh, rOutCh, wg)

	(c.routines[key.key]).Execute(Context(ctx))

	return nil
}

// Send @impl - send to a routine channel
func (c *GlobalContainer) Send(key *RoutineKey, msg RoutineMsg) error {

	if c.inChans[key.key] == nil {
		return fmt.Errorf("Cannot send msg to nil channel for routineKey %#v", key)
	}

	c.inChans[key.key] <- msg
	return nil
}

// Subscribe @impl - subscribe to a given routine channel
func (c *GlobalContainer) Subscribe(key *RoutineKey) (<-chan RoutineMsg, error) {

	var sc chan RoutineMsg
	if c.outChans[key.key] != nil {
		sc = make(chan RoutineMsg)
		c.subChans[key.key] = append(c.subChans[key.key], sc)
	} else {
		return nil, fmt.Errorf("No out channel matching key")
	}

	return sc, nil
}

// Start @impl Continer.Start - Spins up all container routines and msg wiring
func (c *GlobalContainer) Start(wg *sync.WaitGroup) {
	defer (*wg).Done()

	c.containerLock.Lock()

	if c.running {
		err := "Attempting to start a container that is already running is not allowed!"
		c.logger.LogError(err)
		return
	}

	c.running = true
	c.containerLock.Unlock()

	go func(c *GlobalContainer) {
		for k := range c.keys {
			(*wg).Add(1)
			rk := c.keys[k]
			go c.Execute(&rk, wg)
			fmt.Printf("running go execute %s\n", *rk.name)
		}
		wg.Wait()
	}(c)

	for c.running == true {
		for k, oc := range c.outChans {
			subCnt := len(c.subChans)
			if subCnt < 1 {
				continue
			}

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

// IsRunning - impl is container.IsRunning
func (c *GlobalContainer) IsRunning() bool {
	return c.running
}

// Stop @impl Container.Stop - Stops the container routines
func (c *GlobalContainer) Stop() {
	c.running = false
}
