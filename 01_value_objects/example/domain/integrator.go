package domain

type Integrator struct {
	value string
}

func NewIntegrator(value string) (Integrator, error) {
	return Integrator{value}, nil
}

func (i Integrator) Value() string {
	return i.value
}

func (i Integrator) Equals(other Integrator) bool {
	return i.value == other.Value()
}
