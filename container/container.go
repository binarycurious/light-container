package container

import (
	"sync"

	"github.com/binarycurious/light-container/telemetry"
)

// Container - defines a simple container interface for executing tasks in a standard way with access to a global state
type Container interface {
	GetState() *GlobalState                                          //Return the container global state
	GetLogger() telemetry.Logger                                     //Return the container logger
	Subscribe(key *RoutineKey) (<-chan RoutineMsg, error)            //Subscribe to a routine channel
	Send(key *RoutineKey, msg RoutineMsg) error                      //Send msg to a given routine
	Execute(key *RoutineKey) error                                   //Execute a routine
	AddRoutine(routine Routine) RoutineKey                           //Add a routine
	AddNamedRoutine(routineName *string, routine Routine) RoutineKey //Add a routine with a name override
	GetRoutineKey(routineName *string) *RoutineKey                   //Get a routineKey from the string id
	Start(wg *sync.WaitGroup)
	Stop()
	IsRunning() bool
}

// Context - The context ref. exposing limited set of container methods to the ContainerRoutines
type Context interface {
	GetState() *GlobalState
	GetLogger() telemetry.Logger
	GetRoutineName() string
	GetRoutineKey(routineName *string) RoutineKey
	Subscribe(key *RoutineKey) (<-chan RoutineMsg, error)
	Send(key *RoutineKey, msg RoutineMsg) error
	Publish(msg RoutineMsg) error
	GetReceiver() (<-chan RoutineMsg, error)
	ContainerIsRunning() bool
}

// Routine - defines the routine that is executed by the container and the channel that it exports for subscription
type Routine interface {
	Execute(ctx Context) error
	GetName() *string
}

// RoutineMsg - deines the message that is send on routine channels
type RoutineMsg interface {
	GetId() *string
	GetName() *string
	GetMsg() interface{}
}
