package container

import (
	"crypto/sha1"
	"fmt"
	"sync"

	"github.com/binarycurious/light-container/telemetry"
)

// GlobalContainer :  is an implementation of a Global Container
type GlobalContainer struct {
	logger        telemetry.Logger
	state         *GlobalState
	routines      map[string]*ContainerRoutine
	inChans       map[string]chan interface{}
	outChans      map[string][]chan interface{}
	containerLock *sync.Mutex
}

// NewContainer - initialize a new IoC container, nil logger will create a new logger based on the globalstate settings
func (c *GlobalContainer) NewContainer(state GlobalState, logger telemetry.Logger) *Container {

	c.containerLock.Lock()

	if logger == nil {
		logger = telemetry.Logger(&GlobalLogger{hardfail: state.settings.Hardfail})
	}
	c.logger = logger

	validSate := true

	if !state.intitialized {
		msg := "FATAL: Attempt to initialize container state with an un-initialized global state (call NewState())"
		c.logger.LogFatal(&msg)
		validSate = false
	}

	if c.state.intitialized {
		msg := "FATAL: Attempt to set container state after already intialized"
		c.logger.LogFatal(&msg)
		validSate = false
	}

	if validSate {
		c.state = &state
	}

	c.containerLock.Unlock()

	rc := Container(c)
	return &rc
}

// GetRoutineKey : Get a key for given routine name (This does not take into account name conflicts / duplicates)
func (c *GlobalContainer) GetRoutineKey(routineName *string) *routineKey {
	return &routineKey{
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

// AddRoutine : impl of Container.AddRoutine (will modify routineName if there is a conflict)
func (c *GlobalContainer) AddRoutine(routineName *string, routine ContainerRoutine) *routineKey {
	if len(c.routines) == 0 {
		c.routines = make(map[string]*ContainerRoutine, 10)
	}
	rKey := c.GetRoutineKey(routineName)
	_, retry := c.routines[rKey.key]

	retries := 0
	for retry {
		retries++
		*routineName = fmt.Sprintf("%s_%d", *routineName, retries)
		rKey = c.GetRoutineKey(routineName)
		_, retry = c.routines[rKey.key]
	}

	c.routines[rKey.key] = &routine
	return rKey
}

// Execute : impl of container Execute function
func (c *GlobalContainer) Execute(key *routineKey) {
	go (*c.routines[key.key]).Execute(key, Context(c))
}

// Subscribe @impl
func (c *GlobalContainer) Subscribe(key *routineKey) <-chan interface{} {
	panic("not implemented") // TODO: Implement
}

// Send @impl
func (c *GlobalContainer) Send(key *routineKey, msg interface{}) {
	panic("not implemented") // TODO: Implement
}

// Start @impl Continer.Start
func (c *GlobalContainer) Start() {
	panic("not implemented") // TODO: Implement
}

// Stop @impl Container.Stop
func (c *GlobalContainer) Stop() {
	panic("not implemented") // TODO: Implement
}
