package example2

import (
	"errors"
	"regexp"
	"strings"
)

type Phone struct {
	countryCode string
	phoneNumber string
}

var (
	ErrInvalidCountryCode = errors.New("invalid country code")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrEmptyCountryCode   = errors.New("empty country code")
	ErrEmptyPhoneNumber   = errors.New("empty phone number")
)

func NewPhone(countryCode, phoneNumber string) (Phone, error) {
	if strings.TrimSpace(countryCode) == "" {
		return Phone{}, ErrEmptyCountryCode
	}

	if strings.TrimSpace(phoneNumber) == "" {
		return Phone{}, ErrEmptyPhoneNumber
	}

	ok, err := regexp.MatchString(`^[0-9]+$`, countryCode)
	if !ok || err != nil {
		return Phone{}, ErrInvalidCountryCode
	}

	ok, err = regexp.MatchString(`^[0-9]+$`, phoneNumber)
	if !ok || err != nil {
		return Phone{}, ErrInvalidPhoneNumber
	}

	return Phone{
		countryCode: countryCode,
		phoneNumber: phoneNumber,
	}, nil
}

func (p Phone) CountryCode() string {
	return p.countryCode
}

func (p Phone) PhoneNumber() string {
	return p.phoneNumber
}

func (p Phone) EqualTo(other Phone) bool {
	return p.countryCode == other.countryCode && p.phoneNumber == other.phoneNumber
}
