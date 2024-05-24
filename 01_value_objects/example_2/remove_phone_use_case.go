package example2

type (
	RemovePhoneUseCase interface {
		Execute(userID int64, countryCode, phoneNumber string) error
	}

	removePhoneUseCase struct {
		userRepository UserRepository
	}
)

func NewRemovePhoneUseCase(userRepository UserRepository) RemovePhoneUseCase {
	return &removePhoneUseCase{
		userRepository: userRepository,
	}
}

func (u removePhoneUseCase) Execute(userID int64, countryCode, phoneNumber string) error {
	user, err := u.userRepository.Find(userID)
	if err != nil {
		return err
	}

	p, err := NewPhone(countryCode, phoneNumber)
	if err != nil {
		return err
	}

	err = user.RemovePhone(p)
	if err != nil {
		return err
	}

	return u.userRepository.Save(user)
}
