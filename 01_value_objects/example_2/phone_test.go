package example2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPhone(t *testing.T) {
	testCases := []struct {
		name          string
		countryCode   string
		phoneNumber   string
		expectedPhone Phone
		expectedError error
	}{
		{
			name:        "Build phone successfully",
			countryCode: "57",
			phoneNumber: "3103103131",
			expectedPhone: Phone{
				countryCode: "57",
				phoneNumber: "3103103131",
			},
			expectedError: nil,
		},
		{
			name:          "Return error when country code is invalid",
			countryCode:   "adf",
			phoneNumber:   "3103103131",
			expectedPhone: Phone{},
			expectedError: ErrInvalidCountryCode,
		},
		{
			name:          "Return error when phone number is invalid",
			countryCode:   "57",
			phoneNumber:   "adf",
			expectedPhone: Phone{},
			expectedError: ErrInvalidPhoneNumber,
		},
		{
			name:          "Return error when phone number is empty",
			countryCode:   " ",
			phoneNumber:   "31031031",
			expectedPhone: Phone{},
			expectedError: ErrEmptyCountryCode,
		},
		{
			name:          "Return error when phone number is empty",
			countryCode:   "57",
			phoneNumber:   "     ",
			expectedPhone: Phone{},
			expectedError: ErrEmptyPhoneNumber,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := NewPhone(tc.countryCode, tc.phoneNumber)

			assert.Equal(t, tc.expectedPhone, p)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestPhone_CountryCode(t *testing.T) {
	p, err := NewPhone("57", "3103103131")
	assert.NoError(t, err)

	assert.Equal(t, "57", p.CountryCode())
}

func TestPhone_PhoneNumber(t *testing.T) {
	p, err := NewPhone("57", "3103103131")
	assert.NoError(t, err)

	assert.Equal(t, "3103103131", p.PhoneNumber())
}

func TestPhone_EqualTo(t *testing.T) {
	assert := assert.New(t)

	p1, err := NewPhone("57", "3103103131")
	assert.NoError(err)

	p2, err := NewPhone("57", "3103103131")
	assert.NoError(err)

	assert.True(p1.EqualTo(p2))

	p3, err := NewPhone("57", "3103103132")
	assert.NoError(err)

	assert.False(p1.EqualTo(p3))
}
