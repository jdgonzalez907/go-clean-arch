package domain

import (
	"fmt"
	"time"
)

type Changelog struct {
	id         string
	executedBy Integrator
	occurredOn time.Time
	logs       []Log
}

func NewChangelog(id, integrator string, ocurredOn time.Time, logs []Log) (Changelog, error) {
	newIntegrator, err := NewIntegrator(integrator)
	if err != nil {
		return Changelog{}, err
	}

	return Changelog{
		id:         id,
		executedBy: newIntegrator,
		occurredOn: ocurredOn,
		logs:       logs,
	}, nil
}

func (cl Changelog) ExecutedBy() Integrator {
	return cl.executedBy
}

func (cl Changelog) OccurredOn() time.Time {
	return cl.occurredOn
}

func (cl Changelog) Logs() []Log {
	return cl.logs
}

func (cl Changelog) Equals(other Changelog) bool {
	equalsLog := 0

	for _, otherLog := range other.logs {
		for _, log := range cl.logs {
			if otherLog.Equals(log) {
				equalsLog++
			}
		}
	}

	return cl.executedBy.Equals(other.executedBy) &&
		cl.occurredOn == other.occurredOn &&
		equalsLog == len(other.logs)
}

func (cl Changelog) String() string {
	return fmt.Sprintf("Changelog{id: %s, executedBy: %s, occurredOn: %s, logs: %s}", cl.id, cl.executedBy.Value(), cl.occurredOn, cl.logs)
}
