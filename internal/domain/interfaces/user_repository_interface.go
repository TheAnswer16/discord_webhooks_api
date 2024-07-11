package interfaces

type UserRepositoryInterface interface {
	CreateUser(user *User) (*User, error)
	GetUserById(id int) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int) error
}