package dto

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Name     string `json:"name", omitempty:"true"`
	Email    string `json:"email", omitempty:"true"`
	Username string `json:"username", omitempty:"true"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
