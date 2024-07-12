package routes

import (
	"fmt"
	"net/http"

	"github.com/TheAnswer16/discord_webhooks_api/internal/app/controllers"
)

func HandleUsersRoutes(UserController *controllers.UserController) {

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		UserController.CreateUser(w, r)
	})

	http.HandleFunc("/getUserById/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		UserController.GetUserByID(w, r)
	})

	http.HandleFunc("/updateUser/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		UserController.UpdateUser(w, r)
	})

	http.HandleFunc("/deleteUser/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		UserController.DeleteUser(w, r)
	})

}
