package domain

import "fmt"

type Phone struct {
	countryCode string
	phoneNumber string
}

func NewPhone(countryCode string, phoneNumber string) (Phone, error) {
	return Phone{countryCode, phoneNumber}, nil
}

func (p Phone) CountryCode() string {
	return p.countryCode
}

func (p Phone) PhoneNumber() string {
	return p.phoneNumber
}

func (p Phone) FullNumber() string {
	return fmt.Sprintf("(%s) %s", p.countryCode, p.phoneNumber)
}

func (p Phone) Equals(other Phone) bool {
	return p.countryCode == other.CountryCode() && p.phoneNumber == other.PhoneNumber()
}
