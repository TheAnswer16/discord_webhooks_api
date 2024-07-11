package controllers

import (
	
	"encoding/json"
	"net/http"


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
	
	var user *entities.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := uc.UserService.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	user, err := uc.UserService.GetUserByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user *entities.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser, err := uc.UserService.UpdateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	err := uc.UserService.DeleteUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}