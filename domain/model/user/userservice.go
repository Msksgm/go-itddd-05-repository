package user

type UserService struct {
	userRepository UserRepositorier
}

func (us *UserService) NewUserService(userRepository UserRepositorier) error {
	us.userRepository = userRepository
	return nil
}

func (us *UserService) Exists(user *User) (bool, error) {
	user, err := us.userRepository.FindByUserName(user.Name())
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, err
	}
	return true, nil
}
