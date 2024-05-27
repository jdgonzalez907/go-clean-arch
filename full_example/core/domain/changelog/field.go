package changelog

import (
	"fmt"

	"github.com/jdgonzalez907/go-patterns/full_example/core/domain"
)

type Field struct {
	value string
}

const (
	nicknameValue = "nickname"
	phoneValue    = "phone"
)

var (
	supportedFields = []string{
		nicknameValue,
		phoneValue,
	}
)

func NewField(value string) (Field, error) {
	if value == "" {
		return Field{}, domain.ErrEmptyField
	}

	for _, field := range supportedFields {
		if field == value {
			return Field{value: value}, nil
		}
	}

	return Field{}, domain.ErrUnsupportedField
}

func (f Field) Value() string {
	return f.value
}

func (f Field) Equals(other Field) bool {
	return f.value == other.Value()
}

func (f Field) String() string {
	return fmt.Sprintf("Field{value: %s}", f.value)
}

func NicknameField() Field {
	return Field{value: nicknameValue}
}

func PhoneField() Field {
	return Field{value: phoneValue}
}
