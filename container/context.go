package container

import (
	"fmt"

	"github.com/binarycurious/light-container/telemetry"
)

// RoutineContext - the object used for context injection in container routines
type RoutineContext struct {
	key       *RoutineKey
	inChan    <-chan RoutineMsg
	outChan   chan<- RoutineMsg
	container Container
}

// ContainerIsRunning - impl
func (c *RoutineContext) ContainerIsRunning() bool {
	return c.container.IsRunning()
}

// NewRoutineContext - Create a standard context instance
func NewRoutineContext(k *RoutineKey, c Container, inCh <-chan RoutineMsg, outCh chan<- RoutineMsg) *RoutineContext {
	p := RoutineContext{key: k, container: c, inChan: inCh, outChan: outCh}
	return &p
}

// GetRoutineName - Get the name of the current routine
func (c *RoutineContext) GetRoutineName() string {
	return *(c.key).name
}

// Publish - impl of Context publish method (used for publishing messages from a container routine)
func (c *RoutineContext) Publish(msg RoutineMsg) error {
	c.GetLogger().LogDebug(fmt.Sprintf("Publishing message from routine: %s", *c.key.name))
	if c.outChan == nil {
		return fmt.Errorf("Missing out channel in context of routine : (%s)", *c.key.name)
	}
	c.outChan <- msg
	return nil
}

// GetReceiver - context GetReceiver impl
func (c *RoutineContext) GetReceiver() (<-chan RoutineMsg, error) {
	if c.inChan == nil {
		return nil, fmt.Errorf("Missing receiver channel for routine")
	}
	return c.inChan, nil
}

// GetRoutineKey - context GetRoutineKey impl
func (c *RoutineContext) GetRoutineKey(routineName *string) RoutineKey {
	return *c.container.GetRoutineKey(routineName)
}

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
	log := fmt.Sprintf("Subscribed to routine channel: %s , from routine: %s", *key.name, *c.key.name)
	c.GetLogger().LogDebug(log)
	return c.container.Subscribe(key)
}

// Send @impl
func (c *RoutineContext) Send(key *RoutineKey, msg RoutineMsg) error {
	fmt.Print(c.GetRoutineName())
	c.GetLogger().LogDebug(fmt.Sprintf("Sending message %s to routine: (%s) from routine: (%s)", *msg.GetId(), *key.name, *c.key.name))
	return c.container.Send(key, msg)
}
