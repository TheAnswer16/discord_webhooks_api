package services

import (
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/interfaces"
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/entities"
)

type UserService struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUserService(ur interfaces.UserRepositoryInterface) *interfaces.UserServiceInterface {
	return &UserService{
		UserRepository: ur,
	}
}

func (us *UserService) CreateUser(user *entities.User) (*entities.User, error) {
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

