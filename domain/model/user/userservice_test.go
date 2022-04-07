package user

import (
	"fmt"
	"testing"
)

type UserRepositorierStub struct{}

func (us *UserRepositorierStub) FindByUserName(name *UserName) (*User, error) {
	userId, _ := NewUserId("userId")
	userName, _ := NewUserName("userName")
	user, _ := NewUser(*userId, *userName)

	if !userName.Equals(*name) {
		return nil, nil
	}

	return user, nil
}

func (us *UserRepositorierStub) Save(user *User) error {
	return fmt.Errorf("not implemented err")
}

func Test_Exists(t *testing.T) {
	userService := UserService{userRepository: &UserRepositorierStub{}}

	userId, _ := NewUserId("userId")
	userName, _ := NewUserName("userName")
	t.Run("exists", func(t *testing.T) {
		user, _ := NewUser(*userId, *userName)

		isExists, err := userService.Exists(user)
		if err != nil {
			t.Fatal(err)
		}
		if !isExists {
			t.Errorf("isExists must be true but false")
		}
	})
	t.Run("not exists", func(t *testing.T) {
		otherUserName, _ := NewUserName("otherUserName")
		user, _ := NewUser(*userId, *otherUserName)
		isExists, err := userService.Exists(user)
		if err != nil {
			t.Fatal(err)
		}
		if isExists {
			t.Errorf("isExists must be false but true")
		}
	})
}
