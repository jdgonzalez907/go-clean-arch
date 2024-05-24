package example2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	assert := assert.New(t)

	u, err := NewUser(-1, make([]Phone, 0))
	assert.Equal(User{}, u)
	assert.Equal(ErrInvalidID, err)

	u, err = NewUser(1, make([]Phone, 0))
	assert.Equal(User{id: 1, phones: make([]Phone, 0)}, u)
	assert.Nil(err)
}

func TestUser_ID(t *testing.T) {
	assert := assert.New(t)

	u, _ := NewUser(1, make([]Phone, 0))
	assert.Equal(int64(1), u.ID())
}

func TestUser_Phones(t *testing.T) {
	assert := assert.New(t)

	p, _ := NewPhone("57", "3103103131")

	u, _ := NewUser(1, []Phone{p})
	assert.Equal([]Phone{p}, u.Phones())
}

func TestUser_AddPhone(t *testing.T) {
	assert := assert.New(t)

	p, _ := NewPhone("57", "3103103131")
	u, _ := NewUser(1, make([]Phone, 0))

	err := u.AddPhone(p)
	assert.Nil(err)

	err = u.AddPhone(p)
	assert.Equal(ErrPhoneAlreadyExists, err)
}

func TestUser_RemovePhone(t *testing.T) {
	assert := assert.New(t)

	p, _ := NewPhone("57", "3103103131")
	u, _ := NewUser(1, []Phone{p})

	err := u.RemovePhone(p)
	assert.Nil(err)

	err = u.RemovePhone(p)
	assert.Equal(ErrPhoneDoesNotExists, err)
}

func TestUser_UpdatePhone(t *testing.T) {
	assert := assert.New(t)

	p1, _ := NewPhone("57", "3103103131")
	p2, _ := NewPhone("57", "3103103132")
	u, _ := NewUser(1, []Phone{p1})

	err := u.UpdatePhone(p1, p2)
	assert.Nil(err)

	err = u.UpdatePhone(p1, p2)
	assert.Equal(ErrPhoneDoesNotExists, err)
}
