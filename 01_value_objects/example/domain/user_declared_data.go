package domain

import (
	"fmt"
	"time"
)

type UserDeclaredData struct {
	id             int
	brandName      *BrandName
	softDescriptor *SoftDescriptor
	email          *Email
	phone          *Phone
}

func NewUserDeclaredData(id int, brandName, softDescriptor, email, countryCode, phoneNumber *string) (UserDeclaredData, error) {
	var newBrandName *BrandName = nil
	var newSoftDescriptor *SoftDescriptor = nil
	var newEmail *Email = nil
	var newPhone *Phone = nil

	if brandName != nil {
		tmp, err := NewBrandName(*brandName)
		if err != nil {
			return UserDeclaredData{}, err
		}

		newBrandName = &tmp
	}

	if softDescriptor != nil {
		tmp, err := NewSoftDescriptor(*softDescriptor)
		if err != nil {
			return UserDeclaredData{}, err
		}

		newSoftDescriptor = &tmp
	}

	if email != nil {
		tmp, err := NewEmail(*email)
		if err != nil {
			return UserDeclaredData{}, err
		}

		newEmail = &tmp
	}

	if countryCode != nil && phoneNumber != nil {
		tmp, err := NewPhone(*countryCode, *phoneNumber)
		if err != nil {
			return UserDeclaredData{}, err
		}

		newPhone = &tmp
	}

	return UserDeclaredData{
		id:             id,
		brandName:      newBrandName,
		softDescriptor: newSoftDescriptor,
		email:          newEmail,
		phone:          newPhone,
	}, nil
}

func (u UserDeclaredData) ID() int {
	return u.id
}

func (u UserDeclaredData) BrandName() *BrandName {
	return u.brandName
}

func (u UserDeclaredData) SoftDescriptor() *SoftDescriptor {
	return u.softDescriptor
}

func (u UserDeclaredData) Email() *Email {
	return u.email
}

func (u UserDeclaredData) Phone() *Phone {
	return u.phone
}

func (u *UserDeclaredData) Merge(other UserDeclaredData) {
	if other.BrandName() != nil {
		u.brandName = other.BrandName()
	}

	if other.SoftDescriptor() != nil {
		u.softDescriptor = other.SoftDescriptor()
	}

	if other.Email() != nil {
		u.email = other.Email()
	}

	if other.Phone() != nil {
		u.phone = other.Phone()
	}
}

func (u *UserDeclaredData) Update(request DeclaredDataUpdateRequest, idChangelog string, occurredOn time.Time) (Changelog, error) {
	if request.BrandName() == nil &&
		request.SoftDescriptor() == nil &&
		request.Email() == nil &&
		request.Phone() == nil {

		return Changelog{}, nil
	}

	logs := make([]Log, 0)

	if request.BrandName() != nil {
		var old, new *string = nil, nil

		if u.brandName != nil {
			tmp := u.brandName.Value()
			old = &tmp
		}

		if request.BrandName() != nil {
			tmp := request.BrandName().Value()
			new = &tmp
		}

		log, err := NewLog("brand_name", old, new)
		if err != nil {
			return Changelog{}, err
		}

		logs = append(logs, log)

		u.brandName = request.BrandName()
	}

	if request.SoftDescriptor() != nil {
		var old, new *string = nil, nil

		if u.softDescriptor != nil {
			tmp := u.softDescriptor.Value()
			old = &tmp
		}

		if request.SoftDescriptor() != nil {
			tmp := request.SoftDescriptor().Value()
			new = &tmp
		}

		log, err := NewLog("soft_descriptor", old, new)
		if err != nil {
			return Changelog{}, err
		}

		logs = append(logs, log)

		u.softDescriptor = request.SoftDescriptor()
	}

	if request.Email() != nil {
		var old, new *string = nil, nil

		if u.email != nil {
			tmp := u.email.Value()
			old = &tmp
		}

		if request.Email() != nil {
			tmp := request.Email().Value()
			new = &tmp
		}

		log, err := NewLog("email", old, new)
		if err != nil {
			return Changelog{}, err
		}

		logs = append(logs, log)

		u.email = request.Email()
	}

	if request.Phone() != nil {
		var old, new *string = nil, nil

		if u.phone != nil {
			tmp := u.Phone().FullNumber()
			old = &tmp
		}

		if request.Phone() != nil {
			tmp := request.phone.FullNumber()
			new = &tmp
		}

		log, err := NewLog("phone", old, new)
		if err != nil {
			return Changelog{}, err
		}

		logs = append(logs, log)

		u.phone = request.Phone()
	}

	changelog, err := NewChangelog(idChangelog, request.RequestedBy().Value(), occurredOn, logs)
	if err != nil {
		return Changelog{}, err
	}

	return changelog, nil
}

func (u UserDeclaredData) String() string {
	var brandName, softDescriptor, email, phone string

	if u.brandName != nil {
		brandName = u.brandName.Value()
	}

	if u.softDescriptor != nil {
		softDescriptor = u.softDescriptor.Value()
	}

	if u.email != nil {
		email = u.email.Value()
	}

	if u.phone != nil {
		phone = u.Phone().FullNumber()
	}

	return fmt.Sprintf("UserDeclaredData{id: %d, brand_name: %s, soft_descriptor: %s, email: %s, phone: %s}", u.id, brandName, softDescriptor, email, phone)
}
