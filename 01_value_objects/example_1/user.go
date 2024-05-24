package example1

type User struct {
	ID     int64
	Phones []Phone
}

type Phone struct {
	CountryCode string
	PhoneNumber string
}
