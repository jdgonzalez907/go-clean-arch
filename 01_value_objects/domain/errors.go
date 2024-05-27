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
	ErrUserIdMustPositive = errors.New("user id must be positive")
)
