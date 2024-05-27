package changelog

type ChangelogRepository interface {
	Save(changelog Changelog) error
}
