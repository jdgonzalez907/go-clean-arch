package domain

type SoftDescriptor struct {
	value string
}

func NewSoftDescriptor(value string) (SoftDescriptor, error) {
	return SoftDescriptor{value}, nil
}

func (sd SoftDescriptor) Value() string {
	return sd.value
}

func (sd SoftDescriptor) Equals(other SoftDescriptor) bool {
	return sd.value == other.Value()
}
