package example2

import "errors"

type User struct {
	id     int64
	phones []Phone
}

var (
	ErrInvalidID          = errors.New("invalid id")
	ErrPhoneAlreadyExists = errors.New("phone already exists")
	ErrPhoneDoesNotExists = errors.New("phone does not exists")
)

func NewUser(id int64, phones []Phone) (User, error) {
	if id <= 0 {
		return User{}, ErrInvalidID
	}
	return User{
		id:     id,
		phones: phones,
	}, nil
}

func (u *User) ID() int64 {
	return u.id
}

func (u *User) Phones() []Phone {
	return u.phones
}

func (u *User) AddPhone(phone Phone) error {
	for _, p := range u.phones {
		if p.EqualTo(phone) {
			return ErrPhoneAlreadyExists
		}
	}

	u.phones = append(u.phones, phone)

	return nil
}

func (u *User) RemovePhone(phone Phone) error {
	newPhones := make([]Phone, 0)
	exists := false

	for _, p := range u.phones {
		if p.EqualTo(phone) {
			exists = true
			break
		}
		newPhones = append(newPhones, p)
	}

	if !exists {
		return ErrPhoneDoesNotExists
	}

	u.phones = newPhones

	return nil
}

func (u *User) UpdatePhone(oldPhone, newPhone Phone) error {
	for i, p := range u.phones {
		if p.EqualTo(oldPhone) {
			u.phones[i] = newPhone
			return nil
		}
	}

	return ErrPhoneDoesNotExists
}
