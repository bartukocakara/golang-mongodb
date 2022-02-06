package repository

type UserRepository interface {
	SignUp()
	Login()
	GetUsers()
	GetUser(int)
	UpdateUser(int)
	DeleteUser(int)
}