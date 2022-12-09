package container

import (
	"github.com/binarycurious/light-container/config"
	"github.com/binarycurious/light-container/telemetry"
)

type routineKey struct {
	name string
	id   int64
}

// GlobalState for passing global objects around
type GlobalState struct {
	intitialized bool
	settings     config.Settings
}

// NewState - Create global state object
func (gs *GlobalState) NewState(settings *config.Settings) *GlobalState {
	if gs.intitialized {
		lgr := GlobalLogger{hardfail: settings.Hardfail}
		msg := "FATAL: Attempt to set global state after already intialized"
		lgr.LogFatal(&msg)
	}

	gs.settings = *settings
	gs.intitialized = true

	return gs
}

// Container - defines a simple container interface for executing tasks in a standard way with access to a global state
type Container interface {
	GetState() GlobalState
	GetLogger() telemetry.Logger
	AddRoutine(key *string, routine *ContainerRoutine) routineKey
	Execute(key *routineKey)
	Subscribe(key *routineKey) <-chan *interface{}
	SendToChannel(key *routineKey, msg *interface{})
}

// ContainerRoutine - defines the routine that is executed by the container
type ContainerRoutine interface {
	Execute(routineKey *string, container *Container)
}

// GlobalContainer :  is an implementation of a Global Container
type GlobalContainer struct {
	logger   telemetry.Logger
	state    GlobalState
	routines map[string]ContainerRoutine
	inChans  map[string]chan *interface{}
	outChans map[string]chan *interface{}
}
