package domain

import "errors"

var (
	// Phone
	ErrEmptyCountryCode   = errors.New("empty country code")
	ErrEmptyPhoneNumber   = errors.New("empty phone number")
	ErrInvalidCountryCode = errors.New("invalid country code")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")

	// Nickname
	ErrEmptyNickname    = errors.New("empty nickname")
	ErrNicknameTooLong  = errors.New("nickname too long")
	ErrNicknameTooShort = errors.New("nickname too short")
	ErrInvalidNickname  = errors.New("invalid nickname")

	// User
	ErrUserIDMustBePositive                         = errors.New("user id must be positive")
	ErrDeclaredDataUpdateRequestMustHaveInfoToUpate = errors.New("declared data update request must have info to update")
	ErrUserNotFound                                 = errors.New("user not found")

	// Integrator
	ErrInvalidIntegrator = errors.New("invalid integrator")
	ErrEmptyIntegrator   = errors.New("empty integrator")

	// Field
	ErrEmptyField       = errors.New("empty field")
	ErrUnsupportedField = errors.New("unsupported field")

	// Changelog
	ErrChangelogIDMustBePositive = errors.New("changelog id must be positive")
)
