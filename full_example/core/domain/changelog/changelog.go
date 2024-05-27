package changelog

import (
	"fmt"
	"time"

	"github.com/jdgonzalez907/go-patterns/full_example/core/domain"
)

type Changelog struct {
	id         int64
	userID     int64
	executedBy domain.Integrator
	occurredOn time.Time
	logs       []Log
}

func NewChangelog(id, userID int64, executedBy string, occurredOn time.Time) (Changelog, error) {
	if id <= 0 {
		return Changelog{}, domain.ErrChangelogIDMustBePositive
	}

	if userID <= 0 {
		return Changelog{}, domain.ErrUserIDMustBePositive
	}

	newExecutedBy, err := domain.NewIntegrator(executedBy)
	if err != nil {
		return Changelog{}, err
	}

	return Changelog{
		id:         id,
		userID:     userID,
		executedBy: newExecutedBy,
		occurredOn: occurredOn,
	}, nil
}

func (c *Changelog) ID() int64 {
	return c.id
}

func (c *Changelog) UserID() int64 {
	return c.userID
}

func (c *Changelog) ExecutedBy() domain.Integrator {
	return c.executedBy
}

func (c *Changelog) OccurredOn() time.Time {
	return c.occurredOn
}

func (c *Changelog) Logs() []Log {
	return c.logs
}

func (c *Changelog) AddLog(log Log) {
	c.logs = append(c.logs, log)
}

func (c *Changelog) String() string {
	return fmt.Sprintf("Changelog{id: %d, userID: %d, executedBy: %s, occurredOn: %s, logs: %v}", c.id, c.userID, c.executedBy.String(), c.occurredOn, c.logs)
}
