package declareddata

import (
	"fmt"

	"github.com/jdgonzalez907/go-patterns/01_value_objects/domain"
)

type UserDeclaredData struct {
	id       int64
	nickname *Nickname
	phone    *Phone
}

func NewUserDeclaredData(id int64, nickname, country_code, phone_number *string) (UserDeclaredData, error) {
	if id <= 0 {
		return UserDeclaredData{}, domain.ErrUserIdMustPositive
	}

	var newNickname *Nickname = nil
	var newPhone *Phone = nil

	if nickname != nil {
		tmp, err := NewNickname(*nickname)
		if err != nil {
			return UserDeclaredData{}, err
		}

		newNickname = &tmp
	}

	if country_code != nil && phone_number != nil {
		tmp, err := NewPhone(*country_code, *phone_number)
		if err != nil {
			return UserDeclaredData{}, err
		}

		newPhone = &tmp
	}

	return UserDeclaredData{
		id:       id,
		nickname: newNickname,
		phone:    newPhone,
	}, nil
}

func (u *UserDeclaredData) ID() int64 {
	return u.id
}

func (u *UserDeclaredData) Nickname() *Nickname {
	return u.nickname
}

func (u *UserDeclaredData) Phone() *Phone {
	return u.phone
}

func (u *UserDeclaredData) Merge(other UserDeclaredData) {
	if other.Nickname() != nil {
		u.nickname = other.Nickname()
	}

	if other.Phone() != nil {
		u.phone = other.Phone()
	}
}

func (u *UserDeclaredData) String() string {
	var nickname, phone string

	if u.nickname != nil {
		nickname = u.nickname.String()
	}

	if u.phone != nil {
		phone = u.phone.String()
	}

	return fmt.Sprintf("UserDeclaredData{id: %d, nickname: %s, phone: %s}", u.id, nickname, phone)
}
