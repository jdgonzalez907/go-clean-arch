package domain

type UserDeclaredDataRepository interface {
	FindById(id int) (UserDeclaredData, error)
	Save(user UserDeclaredData) error
}
