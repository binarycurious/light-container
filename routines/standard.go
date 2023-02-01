package routines

import (
	"fmt"

	"github.com/binarycurious/light-container/container"
)

// StandardRoutine - impl of a Container Routine for general use
type StandardRoutine struct {
	name    *string
	outChan chan container.RoutineMsg
	routine func(container.Context, chan<- container.RoutineMsg, <-chan container.RoutineMsg) error
}

// NewRoutine - create a new container.Routine for registration
func (sr StandardRoutine) NewRoutine(name string, outChan chan container.RoutineMsg, routine func(container.Context, chan<- container.RoutineMsg, <-chan container.RoutineMsg) error) (container.Routine, error) {
	sr.name = &name
	sr.outChan = outChan
	sr.routine = routine

	return container.Routine(&sr), nil
}

/*container.Routine impls*/

// Execute @impl
func (sr *StandardRoutine) Execute(ctx container.Context, ch chan<- container.RoutineMsg) error {
	return sr.routine(ctx, ch, sr.outChan)
}

// Subscribe @impl
func (sr *StandardRoutine) Subscribe() (<-chan container.RoutineMsg, error) {
	if sr == nil {
		return nil, fmt.Errorf("nil out channel, unable to subscribe")
	}
	return sr.outChan, nil
}

// GetName impl
func (sr *StandardRoutine) GetName() *string {
	return sr.name
}
