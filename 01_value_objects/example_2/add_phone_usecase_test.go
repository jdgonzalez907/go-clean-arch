package example2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitializeAddPhoneUseCase() AddPhoneUseCase {
	inMemoryUserRepository := NewInMemoryUserRepository()

	p, _ := NewPhone("57", "3103103131")
	u, _ := NewUser(1, []Phone{p})

	inMemoryUserRepository.Save(u)

	useCase := NewAddPhoneUseCase(inMemoryUserRepository)

	return useCase
}

func TestAddPhoneUseCase(t *testing.T) {
	validPhone, _ := NewPhone("57", "3103103132")

	testCases := []struct {
		name          string
		userID        int64
		countryCode   string
		phoneNumber   string
		expectedPhone Phone
		expectedError error
	}{
		{
			name:          "Add phone to user successfully",
			userID:        1,
			countryCode:   "57",
			phoneNumber:   "3103103132",
			expectedPhone: validPhone,
			expectedError: nil,
		},
		{
			name:          "Return error when user does not exist",
			userID:        2,
			countryCode:   "57",
			phoneNumber:   "3103103132",
			expectedPhone: Phone{},
			expectedError: ErrUserNotFound,
		},
		{
			name:          "Return error when country code is invalid",
			userID:        1,
			countryCode:   "",
			phoneNumber:   "3103103132",
			expectedPhone: Phone{},
			expectedError: ErrEmptyCountryCode,
		},
		{
			name:          "Return error when phone already exists",
			userID:        1,
			countryCode:   "57",
			phoneNumber:   "3103103131",
			expectedPhone: Phone{},
			expectedError: ErrPhoneAlreadyExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			useCase := InitializeAddPhoneUseCase()

			p, err := useCase.Execute(tc.userID, tc.countryCode, tc.phoneNumber)

			assert.Equal(t, tc.expectedPhone, p)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
