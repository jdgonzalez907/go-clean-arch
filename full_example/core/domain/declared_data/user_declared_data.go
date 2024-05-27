package declareddata

import (
	"fmt"
	"time"

	"github.com/jdgonzalez907/go-patterns/full_example/core/domain"
	"github.com/jdgonzalez907/go-patterns/full_example/core/domain/changelog"
)

type UserDeclaredData struct {
	id            int64
	nickname      *Nickname
	phone         *Phone
	recycledPhone bool
	phoneChanged  bool
}

func NewUserDeclaredData(id int64, nickname, country_code, phone_number *string) (UserDeclaredData, error) {
	if id <= 0 {
		return UserDeclaredData{}, domain.ErrUserIDMustBePositive
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
		id:            id,
		nickname:      newNickname,
		phone:         newPhone,
		recycledPhone: false,
		phoneChanged:  false,
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

func (u *UserDeclaredData) RecycledPhone() bool {
	return u.recycledPhone
}

func (u *UserDeclaredData) PhoneChanged() bool {
	return u.phoneChanged
}

func (u *UserDeclaredData) Merge(other UserDeclaredData) {
	if other.Nickname() != nil {
		u.nickname = other.Nickname()
	}

	if other.Phone() != nil {
		u.phone = other.Phone()
	}

	u.recycledPhone = other.RecycledPhone()
}

func (u *UserDeclaredData) Update(request DeclaredDataUpdateRequest, requestedOn time.Time) (changelog.Changelog, error) {
	if request.Nickname() == nil && request.Phone() == nil {
		return changelog.Changelog{}, domain.ErrDeclaredDataUpdateRequestMustHaveInfoToUpate
	}

	cl, err := changelog.NewChangelog(1, u.id, request.RequestedBy().Value(), requestedOn)
	if err != nil {
		return changelog.Changelog{}, err
	}

	if request.Nickname() != nil {
		nicknameLog, err := u.setNickname(request.Nickname())
		if err != nil {
			return changelog.Changelog{}, err
		}
		cl.AddLog(nicknameLog)
	}

	if request.Phone() != nil {
		log, err := u.setPhone(request.Phone())
		if err != nil {
			return changelog.Changelog{}, err
		}
		cl.AddLog(log)
	}

	return cl, nil
}

func (u *UserDeclaredData) RecyclePhone() {
	u.phone = nil
	u.phoneChanged = true
	u.recycledPhone = true
}

func (u *UserDeclaredData) setNickname(nickname *Nickname) (changelog.Log, error) {
	if u.nickname == nil && nickname == nil ||
		u.nickname != nil && nickname != nil && u.nickname.Equals(*nickname) {
		return changelog.Log{}, nil
	}

	var pointerStrOldNickname, pointerStrNewNickname *string = nil, nil
	if u.nickname != nil {
		tmp := u.nickname.Value()
		pointerStrOldNickname = &tmp
	}

	if nickname != nil {
		tmp := nickname.Value()
		pointerStrNewNickname = &tmp
	}

	log, err := changelog.NewLog(changelog.NicknameField().Value(), pointerStrOldNickname, pointerStrNewNickname)
	if err != nil {
		return changelog.Log{}, err
	}

	u.nickname = nickname

	return log, nil
}

func (u *UserDeclaredData) setPhone(phone *Phone) (changelog.Log, error) {
	if u.phone == nil && phone == nil ||
		u.phone != nil && phone != nil && u.phone.Equals(*phone) {
		return changelog.Log{}, nil
	}

	var pointerStrOldPhone, pointerStrNewPhone *string = nil, nil
	if u.phone != nil {
		tmp := u.phone.FormattedFullNumber()
		pointerStrOldPhone = &tmp
	}

	if phone != nil {
		tmp := phone.FormattedFullNumber()
		pointerStrNewPhone = &tmp
	}

	log, err := changelog.NewLog(changelog.PhoneField().Value(), pointerStrOldPhone, pointerStrNewPhone)
	if err != nil {
		return changelog.Log{}, err
	}

	u.phone = phone
	u.phoneChanged = true

	return log, nil
}

func (u *UserDeclaredData) String() string {
	var nickname, phone string

	if u.nickname != nil {
		nickname = u.nickname.String()
	}

	if u.phone != nil {
		phone = u.phone.String()
	}

	return fmt.Sprintf("UserDeclaredData{id: %d, nickname: %s, phone: %s, recycledPhone: %t, phoneChanged: %t}", u.id, nickname, phone, u.recycledPhone, u.phoneChanged)
}
