package infrastructure

import "github.com/jdgonzalez907/go-patterns/full_example/core/domain/changelog"

type InMemoryChangelogRepository struct {
	database map[int64]changelog.Changelog
}

func NewInMemoryChangelogRepository() changelog.ChangelogRepository {
	return &InMemoryChangelogRepository{
		database: make(map[int64]changelog.Changelog),
	}
}

func (r InMemoryChangelogRepository) Save(changelog changelog.Changelog) error {
	r.database[changelog.ID()] = changelog
	return nil
}
