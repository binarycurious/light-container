package container

import (
	"log"

	"github.com/binarycurious/light-container/telemetry"
)

// NewContainer - initialize a new IoC container
func (c *GlobalContainer) NewContainer(state GlobalState, logger *telemetry.Logger) *Container {
	if c.logger == nil {
		lPtr := telemetry.Logger(&GlobalLogger{hardfail: state.settings.Hardfail})
		logger = &lPtr
	}
	c.logger = *logger

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
		c.state = state
	}
	return Container(c)
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
	c.state = s
}

// Execute : impl of container Execute function
func (c *GlobalContainer) Execute(key string) {
	c.routines[key]
	go c.routines[key].Execute()
}

// AddRoutine : impl of add routine functions for GlobalContainer
func (c *GlobalContainer) AddRoutine(key string, routine *ContainerRoutine) {
	if c.routines == nil {
		c.routines = make(map[string]ContainerRoutine, 5)
	}
	c.routines[key] = *routine
}

// GetState : impl of get state functions for GlobalContainer
func (c *GlobalContainer) GetState() GlobalState {
	return c.state
}
