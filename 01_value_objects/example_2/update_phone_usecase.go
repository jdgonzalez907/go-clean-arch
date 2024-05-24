package example2

type (
	UpdatePhoneUseCase interface {
		Execute(userID int64, oldCountryCode, oldPhoneNumber, newCountryCode, newPhoneNumber string) (Phone, error)
	}

	updatePhoneUseCase struct {
		userRepository UserRepository
	}
)

func NewUpdatePhoneUseCase(userRepository UserRepository) UpdatePhoneUseCase {
	return &updatePhoneUseCase{
		userRepository: userRepository,
	}
}

func (u updatePhoneUseCase) Execute(userID int64, oldCountryCode, oldPhoneNumber, newCountryCode, newPhoneNumber string) (Phone, error) {
	user, err := u.userRepository.Find(userID)
	if err != nil {
		return Phone{}, err
	}

	oldPhone, err := NewPhone(oldCountryCode, oldPhoneNumber)
	if err != nil {
		return Phone{}, err
	}

	newPhone, err := NewPhone(newCountryCode, newPhoneNumber)
	if err != nil {
		return Phone{}, err
	}

	err = user.UpdatePhone(oldPhone, newPhone)
	if err != nil {
		return Phone{}, err
	}

	err = u.userRepository.Save(user)
	if err != nil {
		return Phone{}, err
	}

	return newPhone, nil
}
