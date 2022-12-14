package container

import (
	"crypto/sha1"
	"fmt"
	"sync"

	"github.com/binarycurious/light-container/telemetry"
)

// GlobalContainer :  is an implementation of a Global Container
type GlobalContainer struct {
	logger        *telemetry.Logger
	state         *GlobalState
	routines      map[string]*ContainerRoutine
	inChans       map[string]chan *interface{}
	outChans      map[string][]chan *interface{}
	containerLock sync.Mutex
}

// NewContainer - initialize a new IoC container, nil logger will create a new logger based on the globalstate settings
func (c *GlobalContainer) NewContainer(state GlobalState, logger *telemetry.Logger) *Container {

	c.containerLock.Lock()

	if logger == nil {
		lPtr := telemetry.Logger(&GlobalLogger{hardfail: state.settings.Hardfail})
		logger = &lPtr
	}
	c.logger = logger

	validSate := true

	if !state.intitialized {
		msg := "FATAL: Attempt to initialize container state with an un-initialized global state (call NewState())"
		(*c.logger).LogFatal(&msg)
		validSate = false
	}

	if c.state.intitialized {
		msg := "FATAL: Attempt to set container state after already intialized"
		(*c.logger).LogFatal(&msg)
		validSate = false
	}

	if validSate {
		c.state = &state
	}

	c.containerLock.Unlock()

	rc := Container(c)
	return &rc
}

// GetRoutineKey : Get a key for given routine name (This does not take into account name conflicts / duplicates)
func (c *GlobalContainer) GetRoutineKey(routineName *string) *routineKey {
	return &routineKey{
		name: routineName,
		key:  fmt.Sprintf("%x", sha1.Sum([]byte(*routineName)))}
}

/* Continer Impls */

// GetState @impl
func (c *GlobalContainer) GetState() *GlobalState {
	panic("not implemented") // TODO: Implement
}

// GetLogger @impl
func (c *GlobalContainer) GetLogger() *telemetry.Logger {
	panic("not implemented") // TODO: Implement
}

// AddRoutine @impl
func (c *GlobalContainer) AddRoutine(key *string, routine *ContainerRoutine) *routineKey {
	panic("not implemented") // TODO: Implement
}

// Execute : impl of container Execute function
func (c *GlobalContainer) Execute(key *routineKey) {

	go (*c.routines[key.key]).Execute(key, c)
}

// Subscribe @impl
func (c *GlobalContainer) Subscribe(key *routineKey) <-chan *interface{} {
	panic("not implemented") // TODO: Implement
}

// SendToChannel @impl
func (c *GlobalContainer) Send(key *routineKey, msg *interface{}) {
	panic("not implemented") // TODO: Implement
}
