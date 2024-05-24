package example2

type (
	AddPhoneUseCase interface {
		Execute(userID int64, countryCode, phoneNumber string) (Phone, error)
	}

	addPhoneUseCase struct {
		userRepository UserRepository
	}
)

func NewAddPhoneUseCase(userRepository UserRepository) AddPhoneUseCase {
	return &addPhoneUseCase{
		userRepository: userRepository,
	}
}

func (u addPhoneUseCase) Execute(userID int64, countryCode, phoneNumber string) (Phone, error) {
	user, err := u.userRepository.Find(userID)
	if err != nil {
		return Phone{}, err
	}

	p, err := NewPhone(countryCode, phoneNumber)
	if err != nil {
		return Phone{}, err
	}

	err = user.AddPhone(p)
	if err != nil {
		return Phone{}, err
	}

	err = u.userRepository.Save(user)
	if err != nil {
		return Phone{}, err
	}

	return p, nil
}
