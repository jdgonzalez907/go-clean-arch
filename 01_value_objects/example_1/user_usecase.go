package example1

import (
	"errors"
	"regexp"
)

type (
	UserUseCase interface {
		AddPhone(userId int64, countryCode, phoneNumber string) (*Phone, error)
	}

	userUseCase struct {
		userRepository UserRepository
	}
)

var (
	ErrPhoneAlreadyExists = errors.New("phone already exists")
	ErrInvalidCountryCode = errors.New("invalid country code")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
)

func NewUserUseCase(userRepository UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) AddPhone(userId int64, countryCode, phoneNumber string) (*Phone, error) {
	ok, err := regexp.MatchString(`^[0-9]+$`, countryCode)
	if !ok || err != nil {
		return nil, ErrInvalidCountryCode
	}

	ok, err = regexp.MatchString(`^[0-9]+$`, phoneNumber)
	if !ok || err != nil {
		return nil, ErrInvalidPhoneNumber
	}

	user, err := u.userRepository.Find(userId)
	if err != nil {
		return nil, err
	}

	for _, phone := range user.Phones {
		if phone.CountryCode == countryCode && phone.PhoneNumber == phoneNumber {
			return nil, ErrPhoneAlreadyExists
		}
	}

	p := Phone{CountryCode: countryCode, PhoneNumber: phoneNumber}
	user.Phones = append(user.Phones, p)

	err = u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
