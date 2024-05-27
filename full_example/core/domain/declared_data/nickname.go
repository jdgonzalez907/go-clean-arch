package declareddata

import (
	"fmt"
	"regexp"

	"github.com/jdgonzalez907/go-patterns/full_example/core/domain"
)

type Nickname struct {
	value string
}

const (
	nicknameRegex = `^[a-z0-9]+$`
	minLength     = 3
	maxLength     = 20
)

func NewNickname(value string) (Nickname, error) {
	if value == "" {
		return Nickname{}, domain.ErrEmptyNickname
	}

	if len(value) < minLength {
		return Nickname{}, domain.ErrNicknameTooShort
	}

	if len(value) > maxLength {
		return Nickname{}, domain.ErrNicknameTooLong
	}

	nicknameRegexCompiled, err := regexp.Compile(nicknameRegex)
	if err != nil {
		return Nickname{}, err
	}

	if !nicknameRegexCompiled.MatchString(value) {
		return Nickname{}, domain.ErrInvalidNickname
	}

	return Nickname{value: value}, nil
}

func (n Nickname) Value() string {
	return n.value
}

func (n Nickname) Equals(other Nickname) bool {
	return n.value == other.Value()
}

func (n Nickname) String() string {
	return fmt.Sprintf("Nickname{value: %s}", n.value)
}
