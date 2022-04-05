package user

type UserRepositorier interface {
	FindByUserName(name *UserName) (*User, error)
	// Save(user *User) error
}
