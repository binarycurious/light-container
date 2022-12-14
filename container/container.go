package container

import (
	"github.com/binarycurious/light-container/telemetry"
)

// Container - defines a simple container interface for executing tasks in a standard way with access to a global state
type Container interface {
	GetState() *GlobalState
	GetLogger() telemetry.Logger
	AddRoutine(key *string, routine ContainerRoutine) *routineKey
	Execute(key *routineKey)
	Subscribe(key *routineKey) <-chan interface{}
	Send(key *routineKey, msg interface{})
	GetRoutineKey(key *string) *routineKey
	Start()
	Stop()
}

// Context - The context ref. exposing limited set of container methods to the ContainerRoutines
type Context interface {
	GetState() *GlobalState
	GetLogger() telemetry.Logger
	Subscribe(key *routineKey) <-chan interface{}
	Send(key *routineKey, msg interface{})
}

// ContainerRoutine - defines the routine that is executed by the container
type ContainerRoutine interface {
	Execute(routineKey *routineKey, container Context)
}

// Execute : impl of container Execute function
// func (c *GlobalContainer) Execute(key *routineKey) {
// 	go (*c.routines[key.key]).Execute(*key, &Context(c))
// }

// AddRoutine : impl of add routine functions for GlobalContainer (will modify routineName if there is a conflict)
// func (c *GlobalContainer) AddRoutine(routineName *string, routine *ContainerRoutine) routineKey {
// 	if len(c.routines) == 0 {
// 		c.routines = make(map[string]ContainerRoutine, 10)
// 	}
// 	rKey := c.GetRoutineKey(routineName)
// 	_, retry := c.routines[rKey.key]

// 	retries := 0
// 	for retry {
// 		retries++
// 		*routineName = fmt.Sprintf("%s_%d", *routineName, retries)
// 		rKey = c.GetRoutineKey(routineName)
// 		_, retry = c.routines[rKey.key]
// 	}

// 	c.routines[rKey.key] = *routine
// 	return rKey
// }

// GetState : impl of get state function for GlobalContainer
// func (c *GlobalContainer) GetState() *GlobalState {
// 	return
// }

// // GetLogger : impl of get logger function for GlobalContainer
// func (c *GlobalContainer) GetLogger() *telemetry.Logger {
// 	return &c.logger
// }
