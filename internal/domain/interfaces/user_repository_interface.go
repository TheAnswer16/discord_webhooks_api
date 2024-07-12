package interfaces

import (
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/entities"
)

type UserRepositoryInterface interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(id int) error
}