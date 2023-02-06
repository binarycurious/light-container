package container

// StandardRoutine - impl of a Container Routine for general use
type StandardRoutine struct {
	name    *string
	routine func(Context) error
}

// NewStandardRoutine - create a new Routine for registration
func NewStandardRoutine(name string, routine func(Context) error) (Routine, error) {
	sr := StandardRoutine{}

	sr.name = &name
	sr.routine = routine

	return Routine(&sr), nil
}

/*Routine impls*/

// Execute @impl
func (sr *StandardRoutine) Execute(ctx Context) error {
	return sr.routine(ctx)
}

// GetName impl
func (sr *StandardRoutine) GetName() *string {
	return sr.name
}
