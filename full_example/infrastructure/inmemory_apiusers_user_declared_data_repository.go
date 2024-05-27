package infrastructure

import (
	"github.com/jdgonzalez907/go-patterns/full_example/core/domain"
	declareddata "github.com/jdgonzalez907/go-patterns/full_example/core/domain/declared_data"
)

type InMemoryApiUsersUserDeclaredDataRepository struct {
	database map[int64]declareddata.UserDeclaredData
}

func NewInMemoryApiUsersUserDeclaredDataRepository() declareddata.UserDeclaredDataRepository {
	database := make(map[int64]declareddata.UserDeclaredData)

	strToStrPointer := func(s string) *string {
		return &s
	}

	juan, _ := declareddata.NewUserDeclaredData(1, strToStrPointer("eldevjuanda"), nil, nil)
	andres, _ := declareddata.NewUserDeclaredData(2, strToStrPointer("andresdev"), strToStrPointer("57"), strToStrPointer("3224304262"))
	juana, _ := declareddata.NewUserDeclaredData(3, strToStrPointer("juanita"), strToStrPointer("57"), strToStrPointer("3001102243"))
	andrea, _ := declareddata.NewUserDeclaredData(4, strToStrPointer("andrea"), strToStrPointer("57"), strToStrPointer("3001102243"))

	database[juan.ID()] = juan
	database[andres.ID()] = andres
	database[juana.ID()] = juana
	database[andrea.ID()] = andrea

	return &InMemoryApiUsersUserDeclaredDataRepository{
		database: database,
	}
}

func (r *InMemoryApiUsersUserDeclaredDataRepository) FindByID(id int64) (declareddata.UserDeclaredData, error) {
	user, ok := r.database[id]
	if !ok {
		return declareddata.UserDeclaredData{}, domain.ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryApiUsersUserDeclaredDataRepository) FindByPhone(phone declareddata.Phone) ([]declareddata.UserDeclaredData, error) {
	var users []declareddata.UserDeclaredData
	for _, user := range r.database {
		if user.Phone() != nil && user.Phone().Equals(phone) {
			users = append(users, user)
		}
	}

	return users, nil
}

func (r *InMemoryApiUsersUserDeclaredDataRepository) Save(user declareddata.UserDeclaredData) error {
	r.database[user.ID()] = user
	return nil
}

func (r *InMemoryApiUsersUserDeclaredDataRepository) SaveAll(users []declareddata.UserDeclaredData) error {
	for _, user := range users {
		r.database[user.ID()] = user
	}

	return nil
}
