package services

import (
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/entities"
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/interfaces"
	"github.com/TheAnswer16/discord_webhooks_api/pkg/auth"
)

type UserService struct {
	UserRepository  interfaces.UserRepositoryInterface
	PasswordService *PasswordService
}

func NewUserService(ur interfaces.UserRepositoryInterface, ps *PasswordService) *UserService {
	return &UserService{
		UserRepository:  ur,
		PasswordService: ps,
	}
}

func (us *UserService) CreateUser(user *entities.User) (*entities.User, error) {

	hashedPassword, err := us.PasswordService.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	user.AuthToken = auth.GenerateRandBase64Token(32)
	return us.UserRepository.CreateUser(user)
}

func (us *UserService) GetUserByID(id int) (*entities.User, error) {
	return us.UserRepository.GetUserByID(id)
}

func (us *UserService) UpdateUser(user *entities.User) (*entities.User, error) {
	return us.UserRepository.UpdateUser(user)
}

func (us *UserService) DeleteUser(id int) error {
	return us.UserRepository.DeleteUser(id)
}
