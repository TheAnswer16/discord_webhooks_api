package routes

import (
	"fmt"
	"net/http"
	"gitub.com/TheAnswer16/discord_webhooks_api/internal/app/controllers"
)

func HandleUsersRoutes() {

		http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		})

		http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

			controllers.Register(w, r)
		})
}

