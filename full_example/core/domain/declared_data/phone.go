package declareddata

import (
	"fmt"
	"regexp"

	"github.com/jdgonzalez907/go-patterns/full_example/core/domain"
)

type Phone struct {
	countryCode string
	phoneNumber string
}

const (
	countryCodeRegex = `^[1-9][0-9]{0,2}$`
	phoneNumberRegex = `^[0-9]{7,15}$`
)

func NewPhone(countryCode string, phoneNumber string) (Phone, error) {
	if countryCode == "" {
		return Phone{}, domain.ErrEmptyCountryCode
	}
	if phoneNumber == "" {
		return Phone{}, domain.ErrEmptyPhoneNumber
	}

	countryCodeRegexCompiled, err := regexp.Compile(countryCodeRegex)
	if err != nil {
		return Phone{}, err
	}
	phoneNumberRegexCompiled, err := regexp.Compile(phoneNumberRegex)
	if err != nil {
		return Phone{}, err
	}

	if !countryCodeRegexCompiled.MatchString(countryCode) {
		return Phone{}, domain.ErrInvalidCountryCode
	}
	if !phoneNumberRegexCompiled.MatchString(phoneNumber) {
		return Phone{}, domain.ErrInvalidPhoneNumber
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

func (p Phone) Equals(other Phone) bool {
	return p.countryCode == other.CountryCode() && p.phoneNumber == other.PhoneNumber()
}

func (p Phone) FormattedFullNumber() string {
	return fmt.Sprintf("+%s %s", p.countryCode, p.phoneNumber)
}

func (p Phone) String() string {
	return fmt.Sprintf("Phone{countryCode: %s, phoneNumber: %s}", p.countryCode, p.phoneNumber)
}
