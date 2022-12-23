package container

import "github.com/binarycurious/light-container/telemetry"

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

// Publish - impl of Context publish method (used for publishing messages from a container routine)
func (c *RoutineContext) Publish(msg interface{}) error {
	panic("to implement")
}

func (c *RoutineContext) GetState() *GlobalState {
	return c.container.GetState()
}

func (c *RoutineContext) GetLogger() telemetry.Logger {
	return c.container.GetLogger()
}

func (c *RoutineContext) Subscribe(key *RoutineKey) <-chan interface{} {
	return c.container.Subscribe(key)
}

func (c *RoutineContext) Send(key *RoutineKey, msg interface{}) error {
	return c.container.Send(key, msg)
}
