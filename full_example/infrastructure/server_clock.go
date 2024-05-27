package infrastructure

import (
	"time"

	"github.com/jdgonzalez907/go-patterns/full_example/core/domain/shared"
)

type ServerClock struct {
}

func NewServerClock() shared.Clock {
	return &ServerClock{}
}

func (s *ServerClock) Now() time.Time {
	return time.Now()
}
