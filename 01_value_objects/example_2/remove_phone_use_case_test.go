package example2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitializeRemovePhoneUseCase() RemovePhoneUseCase {
	inMemoryUserRepository := NewInMemoryUserRepository()

	p, _ := NewPhone("57", "3103103131")
	u, _ := NewUser(1, []Phone{p})

	inMemoryUserRepository.Save(u)

	useCase := NewRemovePhoneUseCase(inMemoryUserRepository)

	return useCase
}

func TestRemovePhoneUseCase(t *testing.T) {
	testCases := []struct {
		name          string
		userID        int64
		countryCode   string
		phoneNumber   string
		expectedError error
	}{
		{
			name:          "Remove phone successfully",
			userID:        1,
			countryCode:   "57",
			phoneNumber:   "3103103131",
			expectedError: nil,
		},
		{
			name:          "Return error when user not exists",
			userID:        2,
			countryCode:   "57",
			phoneNumber:   "3103103131",
			expectedError: ErrUserNotFound,
		},
		{
			name:          "Return error when phone number is empty",
			userID:        1,
			countryCode:   "57",
			phoneNumber:   "",
			expectedError: ErrEmptyPhoneNumber,
		},
		{
			name:          "Return error when phone not exists",
			userID:        1,
			countryCode:   "58",
			phoneNumber:   "3103103131",
			expectedError: ErrPhoneDoesNotExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			useCase := InitializeRemovePhoneUseCase()

			err := useCase.Execute(tc.userID, tc.countryCode, tc.phoneNumber)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
