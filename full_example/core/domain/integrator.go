package domain

import (
	"fmt"
	"regexp"
)

type Integrator struct {
	value string
}

const (
	integratorRegex = `^[a-z0-9_-]+$`
)

func NewIntegrator(value string) (Integrator, error) {
	if value == "" {
		return Integrator{}, ErrEmptyIntegrator
	}

	integratorRegexCompiled, err := regexp.Compile(integratorRegex)
	if err != nil {
		return Integrator{}, err
	}

	if !integratorRegexCompiled.MatchString(value) {
		return Integrator{}, ErrInvalidIntegrator
	}

	return Integrator{value: value}, nil
}

func (i Integrator) Value() string {
	return i.value
}

func (i Integrator) Equals(other Integrator) bool {
	return i.value == other.Value()
}

func (i Integrator) String() string {
	return fmt.Sprintf("Integrator{value: %s}", i.value)
}
