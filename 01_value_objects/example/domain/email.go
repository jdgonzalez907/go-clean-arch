package domain

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	return Email{value}, nil
}

func (e Email) Value() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return e.value == other.Value()
}
