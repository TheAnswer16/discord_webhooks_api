package repository

import (
	"database/sql"
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/interfaces"
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(user *entities.User) (*User, error) {
	
	var user *entities.User

	query := "INSERT INTO users (name, username, password, email, auth_token) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, username, email, auth_token"

	err := ur.DB.QueryRow(query, user.Name, user.Username, user.Password, user.Email, user.AuthToken).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.AuthToken)

	if err != nil {
		return err
	}

	return user, nil
}