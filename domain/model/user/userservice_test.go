package user

import "testing"

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

func Test_Exists(t *testing.T) {
	userService := UserService{userRepository: &UserRepositorierStub{}}

	userId, _ := NewUserId("userId")
	userName, _ := NewUserName("userName")
	user, _ := NewUser(*userId, *userName)

	isExists, err := userService.Exists(user)
	if err != nil {
		t.Fatal(err)
	}
	if !isExists {
		t.Errorf("isExists must be %v but %v", isExists, isExists)
	}
}
