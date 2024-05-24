package example2

import "errors"

type (
	UserRepository interface {
		Find(id int64) (User, error)
		Save(user User) error
	}

	inMemoryUserRepository struct {
		users map[int64]User
	}
)

var ErrUserNotFound = errors.New("user not found")

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make(map[int64]User),
	}
}

func (r *inMemoryUserRepository) Find(id int64) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (r *inMemoryUserRepository) Save(user User) error {
	r.users[user.ID()] = user
	return nil
}
