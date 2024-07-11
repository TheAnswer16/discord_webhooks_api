package server

import (
	"fmt"
	"net/http"
	"os"
	"github.com/TheAnswer16/discord_webhooks_api/internal/app/routes"
)

func SetupServer() {


	fmt.Println("Server starting on port:", os.Getenv("HTTP_PORT"))

	err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), nil)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	
}