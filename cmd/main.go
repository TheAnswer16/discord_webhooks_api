package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/TheAnswer16/discord_webhooks_api/internal/infra/server"
	"os"
)

func main() {
	
	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load()

		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}

	server.SetupServer()

}