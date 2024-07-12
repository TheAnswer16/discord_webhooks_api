package repository

import (
	"database/sql"

	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/entities"
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/interfaces"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(user *entities.User) (*entities.User, error) {

	query := "INSERT INTO users (name, username, password, email, auth_token) VALUES ($1, $2, $3, $4, $5) RETURNING id_users, name, username, email, auth_token"

	err := ur.DB.QueryRow(query, user.Name, user.Username, user.Password, user.Email, user.AuthToken).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.AuthToken)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(id int) (*entities.User, error) {

	var user *entities.User

	query := "SELECT id, name, username, email, auth_token FROM users WHERE id = $1"

	err := ur.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.AuthToken)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(user *entities.User) (*entities.User, error) {

	query := "UPDATE users SET name = $1, username = $2, email = $3 WHERE id = $4 RETURNING id, name, username, email, auth_token"

	err := ur.DB.QueryRow(query, user.Name, user.Username, user.Email, user.ID).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.AuthToken)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) DeleteUser(id int) error {

	query := "DELETE FROM users WHERE id = $1"

	_, err := ur.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
