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
