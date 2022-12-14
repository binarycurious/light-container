package container

import (
	"crypto/sha1"
	"fmt"
	"log"
	"sync"

	"github.com/binarycurious/light-container/telemetry"
)

// Container - defines a simple container interface for executing tasks in a standard way with access to a global state
type Container interface {
	GetState() *GlobalState
	GetLogger() *telemetry.Logger
	AddRoutine(key *string, routine *ContainerRoutine) *routineKey
	Execute(key *routineKey)
	Subscribe(key *routineKey) <-chan *interface{}
	SendToChannel(key *routineKey, msg *interface{})
	GetRoutineKey(key *string) *routineKey
}

// Context - The context ref. exposing limited set of container methods to the ContainerRoutines
type Context interface {
	GetState() *GlobalState
	GetLogger() *telemetry.Logger
	Subscribe(key *routineKey) <-chan *interface{}
	SendToChannel(key *routineKey, msg *interface{})
}

// ContainerRoutine - defines the routine that is executed by the container
type ContainerRoutine interface {
	Execute(routineKey routineKey, container *Context)
}

// GlobalContainer :  is an implementation of a Global Container
type GlobalContainer struct {
	logger        *telemetry.Logger
	state         *GlobalState
	routines      *map[string]*ContainerRoutine
	inChans       *map[string]chan *interface{}
	outChans      *map[string]chan *interface{}
	containerLock sync.Mutex
}

// NewContainer - initialize a new IoC container, nil logger will create a new logger based on the globalstate settings
func (c *GlobalContainer) NewContainer(state GlobalState, logger *telemetry.Logger) *Container {

	c.containerLock.Lock()

	if logger == nil {
		lPtr := telemetry.Logger(&GlobalLogger{hardfail: state.settings.Hardfail})
		logger = &lPtr
	}
	c.logger = logger

	validSate := true

	if !state.intitialized {
		msg := "FATAL: Attempt to initialize container state with an un-initialized global state (call NewState())"
		(*c.logger).LogFatal(&msg)
		validSate = false
	}

	if c.state.intitialized {
		msg := "FATAL: Attempt to set container state after already intialized"
		(*c.logger).LogFatal(&msg)
		validSate = false
	}

	if validSate {
		c.state = &state
	}

	c.containerLock.Unlock()

	rc := Container(*c)
	return &rc
}

// SetState set and initialize the global state immutably (can only be called on a container with uninitialized state)
func (c *GlobalContainer) SetState(s GlobalState) {

	if c.state.intitialized {
		failMsg := "FATAL: Attempt to set container state after already intialized"
		if !s.settings.Hardfail {
			log.Print(failMsg)
			return
		}
		log.Fatal(failMsg)
	}
	if !s.intitialized {
		log.Fatal("FATAL: Attempt to initialize container state with an un-initialized global state (call NewState())")
	}
	c.state = *s
}

// GetRoutineKey : Get a key for given routine name (This does not take into account name conflicts / duplicates)
func (c *GlobalContainer) GetRoutineKey(routineName *string) routineKey {
	return routineKey{
		name: routineName,
		key:  fmt.Sprintf("%x", sha1.Sum([]byte(*routineName)))}
}

// Execute : impl of container Execute function
func (c *GlobalContainer) Execute(key *routineKey) {
	go (*c.routines[key.key]).Execute(*key, c)
}

// AddRoutine : impl of add routine functions for GlobalContainer (will modify routineName if there is a conflict)
func (c *GlobalContainer) AddRoutine(routineName *string, routine *ContainerRoutine) routineKey {
	if len(c.routines) == 0 {
		c.routines = make(map[string]ContainerRoutine, 10)
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

	c.routines[rKey.key] = *routine
	return rKey
}

// GetState : impl of get state function for GlobalContainer
func (c *GlobalContainer) GetState() *GlobalState {
	return
}

// GetLogger : impl of get logger function for GlobalContainer
func (c *GlobalContainer) GetLogger() *telemetry.Logger {
	return &c.logger
}
