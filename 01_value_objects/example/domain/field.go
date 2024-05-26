package domain

type Field struct {
	value string
}

func NewField(value string) (Field, error) {
	return Field{value}, nil
}

func (f Field) Value() string {
	return f.value
}

func (f Field) Equals(other Field) bool {
	return f.value == other.Value()
}
