package container

import (
	"fmt"

	"github.com/binarycurious/light-container/telemetry"
)

// RoutineContext - the object used for context injection in container routines
type RoutineContext struct {
	key       *RoutineKey
	container Container
}

// NewRoutineContext - Create a standard context instance
func (RoutineContext) NewRoutineContext(k *RoutineKey, c Container) *RoutineContext {
	p := RoutineContext{key: k}
	return &p
}

// // Publish - impl of Context publish method (used for publishing messages from a container routine)
// func (c *RoutineContext) Publish(msg RoutineMsg) error {
// 	log := fmt.Sprintf("Publishing message from routine: %#v", c.key)
// 	c.GetLogger().LogDebug(&log)
// 	return c.container.Publish(c.key, msg)
// }

// GetState @impl
func (c *RoutineContext) GetState() *GlobalState {
	return c.container.GetState()
}

// GetLogger @impl
func (c *RoutineContext) GetLogger() telemetry.Logger {
	return c.container.GetLogger()
}

// Subscribe @impl
func (c *RoutineContext) Subscribe(key *RoutineKey) (<-chan RoutineMsg, error) {
	log := fmt.Sprintf("Subscribed to routine channel: %#v , from routine: %#v", key, c.key)
	c.GetLogger().LogDebug(&log)
	return c.container.Subscribe(key)
}

// Send @impl
func (c *RoutineContext) Send(key *RoutineKey, msg RoutineMsg) error {
	log := fmt.Sprintf("Sending message to routine: %#v , from routine: %#v", key, c.key)
	c.GetLogger().LogDebug(&log)
	return c.container.Send(key, msg)
}
