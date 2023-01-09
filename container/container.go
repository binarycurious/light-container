package container

import (
	"github.com/binarycurious/light-container/telemetry"
)

// Container - defines a simple container interface for executing tasks in a standard way with access to a global state
type Container interface {
	GetState() *GlobalState                              //Return the container global state
	GetLogger() telemetry.Logger                         //Return the container logger
	Subscribe(key *RoutineKey) <-chan interface{}        //Subscribe to a routine channel
	Send(key *RoutineKey, msg interface{}) error         //Send to a routine
	Publish(key *RoutineKey, msg interface{}) error      //Publish from a routine
	Execute(key *RoutineKey)                             //Execute a routine
	AddRoutine(key *string, routine Routine) *RoutineKey //Add a routine
	GetRoutineKey(key *string) *RoutineKey               //Get a routineKey from the string id
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

// Routine - defines the routine that is executed by the container and the channel that it exports for subscription
type Routine interface {
	Execute(ctx Context) <-chan interface{}
}
