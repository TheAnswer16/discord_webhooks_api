package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TheAnswer16/discord_webhooks_api/internal/app/controllers"
	"github.com/TheAnswer16/discord_webhooks_api/internal/app/repository"
	"github.com/TheAnswer16/discord_webhooks_api/internal/app/routes"
	"github.com/TheAnswer16/discord_webhooks_api/internal/app/services"
	"github.com/TheAnswer16/discord_webhooks_api/internal/infra/server"
	"github.com/TheAnswer16/discord_webhooks_api/pkg/database"
	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load()

		if err != nil {
			fmt.Println("Error loading .env file")
		}
	}

	db, err := database.NewConnection("user=postgres password=@Alex1340010ai16jk dbname=discord_webhook_management sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	routes.HandleUsersRoutes(userController)

	server.SetupServer()

}
