package declareddata

import (
	"fmt"
	"time"

	"github.com/jdgonzalez907/go-patterns/full_example/core/domain"
)

type DeclaredDataUpdateRequest struct {
	nickname    *Nickname
	phone       *Phone
	requestedBy domain.Integrator
	requestedOn time.Time
}

func NewDeclaredDataUpdateRequest(nickname, country_code, phone_number *string, requestedBy string, requestedOn time.Time) (DeclaredDataUpdateRequest, error) {
	var newNickname *Nickname = nil
	var newPhone *Phone = nil
	var err error = nil

	if nickname != nil {
		nickname, err := NewNickname(*nickname)
		if err != nil {
			return DeclaredDataUpdateRequest{}, err
		}

		newNickname = &nickname
	}

	if country_code != nil && phone_number != nil {
		phone, err := NewPhone(*country_code, *phone_number)
		if err != nil {
			return DeclaredDataUpdateRequest{}, err
		}

		newPhone = &phone
	}

	newRequestedBy, err := domain.NewIntegrator(requestedBy)
	if err != nil {
		return DeclaredDataUpdateRequest{}, err
	}

	return DeclaredDataUpdateRequest{
		nickname:    newNickname,
		phone:       newPhone,
		requestedBy: newRequestedBy,
		requestedOn: requestedOn,
	}, nil
}

func (r DeclaredDataUpdateRequest) Nickname() *Nickname {
	return r.nickname
}

func (r DeclaredDataUpdateRequest) Phone() *Phone {
	return r.phone
}

func (r DeclaredDataUpdateRequest) RequestedBy() domain.Integrator {
	return r.requestedBy
}

func (r DeclaredDataUpdateRequest) RequestedOn() time.Time {
	return r.requestedOn
}

func (r DeclaredDataUpdateRequest) String() string {
	var nickname, phone string

	if r.nickname != nil {
		nickname = r.nickname.String()
	}

	if r.phone != nil {
		phone = r.phone.String()
	}

	return fmt.Sprintf("DeclaredDataUpdateRequest{nickname: %s, phone: %s, requestedBy: %s, requestedOn: %s}", nickname, phone, r.requestedBy.String(), r.requestedOn)
}
