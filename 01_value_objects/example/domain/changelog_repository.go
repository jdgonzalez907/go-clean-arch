package domain

type ChangelogRepository interface {
	Save(changelog Changelog) error
}
