package routines

import (
	"github.com/binarycurious/light-container/container"
)

// StandardRoutine - impl of a Container Routine for general use
type StandardRoutine struct {
	name    *string
	routine func(container.Context) error
}

// NewStandardRoutine - create a new container.Routine for registration
func NewStandardRoutine(name string, routine func(container.Context) error) (container.Routine, error) {
	sr := StandardRoutine{}

	sr.name = &name
	sr.routine = routine

	return container.Routine(&sr), nil
}

/*container.Routine impls*/

// Execute @impl
func (sr *StandardRoutine) Execute(ctx container.Context) error {
	return sr.routine(ctx)
}

// GetName impl
func (sr *StandardRoutine) GetName() *string {
	return sr.name
}
