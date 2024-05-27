package declareddata

type UserDeclaredDataRepository interface {
	FindByID(id int64) (UserDeclaredData, error)
	FindByPhone(phone Phone) ([]UserDeclaredData, error)
	Save(user UserDeclaredData) error
	SaveAll(users []UserDeclaredData) error
}
