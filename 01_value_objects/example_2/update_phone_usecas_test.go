package example2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitializeUpdatePhoneUseCase() UpdatePhoneUseCase {
	inMemoryUserRepository := NewInMemoryUserRepository()

	p, _ := NewPhone("57", "3103103131")
	u, _ := NewUser(1, []Phone{p})

	inMemoryUserRepository.Save(u)

	useCase := NewUpdatePhoneUseCase(inMemoryUserRepository)

	return useCase
}

func TestUpdatePhoneUseCase(t *testing.T) {
	validPhone, _ := NewPhone("55", "999999999")

	testCases := []struct {
		name           string
		userID         int64
		oldCountryCode string
		oldPhoneNumber string
		newCountryCode string
		newPhoneNumber string
		expectedPhone  Phone
		expectedError  error
	}{
		{
			name:           "Update phone successfully",
			userID:         1,
			oldCountryCode: "57",
			oldPhoneNumber: "3103103131",
			newCountryCode: "55",
			newPhoneNumber: "999999999",
			expectedPhone:  validPhone,
			expectedError:  nil,
		},
		{
			name:           "Return error when old phone number is invalid",
			userID:         1,
			oldCountryCode: "57",
			oldPhoneNumber: "fdsafsa",
			newCountryCode: "55",
			newPhoneNumber: "999999999",
			expectedPhone:  Phone{},
			expectedError:  ErrInvalidPhoneNumber,
		},
		{
			name:           "Return error when new country code is empty",
			userID:         1,
			oldCountryCode: "",
			oldPhoneNumber: "3103103131",
			newCountryCode: "55",
			newPhoneNumber: "999999999",
			expectedPhone:  Phone{},
			expectedError:  ErrEmptyCountryCode,
		},
		{
			name:           "Update phone successfully",
			userID:         1,
			oldCountryCode: "57",
			oldPhoneNumber: "3103103132",
			newCountryCode: "55",
			newPhoneNumber: "999999999",
			expectedPhone:  Phone{},
			expectedError:  ErrPhoneDoesNotExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			useCase := InitializeUpdatePhoneUseCase()

			p, err := useCase.Execute(tc.userID, tc.oldCountryCode, tc.oldPhoneNumber, tc.newCountryCode, tc.newPhoneNumber)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedPhone, p)
		})
	}
}
