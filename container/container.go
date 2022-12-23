package container

import (
	"github.com/binarycurious/light-container/telemetry"
)

// Container - defines a simple container interface for executing tasks in a standard way with access to a global state
type Container interface {
	GetState() *GlobalState
	GetLogger() telemetry.Logger
	Subscribe(key *RoutineKey) <-chan interface{}
	Send(key *RoutineKey, msg interface{}) error

	Execute(key *RoutineKey)
	AddRoutine(key *string, routine Routine) *RoutineKey
	GetRoutineKey(key *string) *RoutineKey
	Start()
	Stop()
}

// Context - The context ref. exposing limited set of container methods to the ContainerRoutines
type Context interface {
	GetState() *GlobalState
	GetLogger() telemetry.Logger
	Subscribe(key *RoutineKey) <-chan interface{}
	Send(key *RoutineKey, msg interface{}) error
	Publish(msg interface{}) error
}

// Routine - defines the routine that is executed by the container
type Routine interface {
	Execute(ctx Context) <-chan interface{}
}
