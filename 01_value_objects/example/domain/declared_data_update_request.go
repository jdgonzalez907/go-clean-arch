package domain

import (
	"fmt"
	"time"
)

type DeclaredDataUpdateRequest struct {
	brandName      *BrandName
	softDescriptor *SoftDescriptor
	email          *Email
	phone          *Phone
	requestedBy    Integrator
	requestedOn    time.Time
}

func NewDeclaredDataUpdateRequest(brandName, softDescriptor, email, countryCode, phoneNumber *string, requestedBy string, requestedOn time.Time) (DeclaredDataUpdateRequest, error) {
	var newBrandName *BrandName
	if brandName != nil {
		tmp, err := NewBrandName(*brandName)
		if err != nil {
			return DeclaredDataUpdateRequest{}, err
		}
		newBrandName = &tmp
	}

	var newSoftDescriptor *SoftDescriptor
	if softDescriptor != nil {
		tmp, err := NewSoftDescriptor(*softDescriptor)
		if err != nil {
			return DeclaredDataUpdateRequest{}, err
		}
		newSoftDescriptor = &tmp
	}

	var newEmail *Email
	if email != nil {
		tmp, err := NewEmail(*email)
		if err != nil {
			return DeclaredDataUpdateRequest{}, err
		}
		newEmail = &tmp
	}

	var newPhone *Phone
	if countryCode != nil && phoneNumber != nil {
		tmp, err := NewPhone(*countryCode, *phoneNumber)
		if err != nil {
			return DeclaredDataUpdateRequest{}, err
		}
		newPhone = &tmp
	}

	newIntegrator, err := NewIntegrator(requestedBy)
	if err != nil {
		return DeclaredDataUpdateRequest{}, err
	}

	return DeclaredDataUpdateRequest{
		brandName:      newBrandName,
		softDescriptor: newSoftDescriptor,
		email:          newEmail,
		phone:          newPhone,
		requestedBy:    newIntegrator,
		requestedOn:    requestedOn,
	}, nil
}

func (r DeclaredDataUpdateRequest) BrandName() *BrandName {
	return r.brandName
}

func (r DeclaredDataUpdateRequest) SoftDescriptor() *SoftDescriptor {
	return r.softDescriptor
}

func (r DeclaredDataUpdateRequest) Email() *Email {
	return r.email
}

func (r DeclaredDataUpdateRequest) Phone() *Phone {
	return r.phone
}

func (r DeclaredDataUpdateRequest) RequestedBy() Integrator {
	return r.requestedBy
}

func (r DeclaredDataUpdateRequest) RequestedOn() time.Time {
	return r.requestedOn
}

func (r DeclaredDataUpdateRequest) String() string {
	var brandName, softDescriptor, email, phone string

	if r.brandName != nil {
		brandName = r.brandName.Value()
	}

	if r.softDescriptor != nil {
		softDescriptor = r.softDescriptor.Value()
	}

	if r.email != nil {
		email = r.email.Value()
	}

	if r.phone != nil {
		phone = r.phone.FullNumber()
	}

	return fmt.Sprintf("DeclaredDataUpdateRequest{brandName: %s, softDescriptor: %s, email: %s, phone: %s, requestedBy: %s, requestedOn: %s}", brandName, softDescriptor, email, phone, r.requestedBy.Value(), r.requestedOn)
}
