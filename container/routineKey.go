package container

type RoutineKey struct {
	name *string
	key  string
}

// GetKey - return the routine key
func (rk *RoutineKey) GetKey() string {
	return rk.key
}
