package domain

type BrandName struct {
	value string
}

func NewBrandName(value string) (BrandName, error) {
	return BrandName{value: value}, nil
}

func (b BrandName) Value() string {
	return b.value
}

func (b BrandName) Equals(other BrandName) bool {
	return b.value == other.Value()
}
