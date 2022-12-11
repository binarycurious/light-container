package container

import (
	"github.com/binarycurious/light-container/config"
)

type routineKey struct {
	name *string
	key  string
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
