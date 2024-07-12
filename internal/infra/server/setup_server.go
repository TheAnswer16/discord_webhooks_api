package server

import (
	"fmt"
	"net/http"
	"os"
)

func SetupServer() {

	fmt.Println("Server starting on port:", os.Getenv("HTTP_PORT"))

	err := http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}
