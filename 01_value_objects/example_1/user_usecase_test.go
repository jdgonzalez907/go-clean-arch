package example1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Initialize() UserUseCase {
	inMemoryUserRepository := NewInMemoryUserRepository()
	inMemoryUserRepository.Save(User{
		ID:     1,
		Phones: []Phone{{CountryCode: "57", PhoneNumber: "3103103131"}},
	})

	useCase := NewUserUseCase(inMemoryUserRepository)

	return useCase
}

func TestAddPhone(t *testing.T) {
	testCases := []struct {
		name          string
		userID        int64
		countryCode   string
		phoneNumber   string
		expectedPhone *Phone
		expectedError error
	}{
		{
			name:          "Add successfully",
			userID:        1,
			countryCode:   "57",
			phoneNumber:   "3103103132",
			expectedPhone: &Phone{CountryCode: "57", PhoneNumber: "3103103132"},
			expectedError: nil,
		},
		{
			name:          "Return error when user does not exist",
			userID:        2,
			countryCode:   "57",
			phoneNumber:   "3103103131",
			expectedPhone: nil,
			expectedError: ErrUserNotFound,
		},
		{
			name:          "Return error when user has same phone",
			userID:        1,
			countryCode:   "57",
			phoneNumber:   "3103103131",
			expectedPhone: nil,
			expectedError: ErrPhoneAlreadyExists,
		},
		{
			name:          "Return error when country code is invalid",
			userID:        1,
			countryCode:   "321 Invalid",
			phoneNumber:   "3103103131",
			expectedPhone: nil,
			expectedError: ErrInvalidCountryCode,
		},
		{
			name:          "Return error when phone number is invalid",
			userID:        1,
			countryCode:   "57",
			phoneNumber:   "Invalid 3123",
			expectedPhone: nil,
			expectedError: ErrInvalidPhoneNumber,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			useCase := Initialize()

			p, err := useCase.AddPhone(tc.userID, tc.countryCode, tc.phoneNumber)

			assert.Equal(t, tc.expectedPhone, p)

			if tc.expectedError != nil {
				assert.ErrorIs(t, tc.expectedError, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
