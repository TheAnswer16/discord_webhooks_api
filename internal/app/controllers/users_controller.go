package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TheAnswer16/discord_webhooks_api/internal/app/dto"
	"github.com/TheAnswer16/discord_webhooks_api/internal/app/services"
	"github.com/TheAnswer16/discord_webhooks_api/internal/domain/entities"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		UserService: us,
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var createUserDTO *dto.CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&createUserDTO)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &entities.User{
		Name:     createUserDTO.Name,
		Username: createUserDTO.Username,
		Password: createUserDTO.Password,
		Email:    createUserDTO.Email,
	}

	createdUser, err := uc.UserService.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Path[len("/getUserById/"):]

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uc.UserService.GetUserByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Path[len("/updateUser/"):]

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updateUserDTO dto.UpdateUserDTO

	err = json.NewDecoder(r.Body).Decode(&updateUserDTO)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &entities.User{
		ID:       id,
		Name:     updateUserDTO.Name,
		Email:    updateUserDTO.Email,
		Username: updateUserDTO.Username,
	}

	updatedUser, err := uc.UserService.UpdateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Path[len("/deleteUser/"):]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = uc.UserService.DeleteUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
