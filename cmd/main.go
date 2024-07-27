package main

import (
	"log"
	"net/http"

	"github.com/bigpandaboy2/project-management-service/internal/app/config"
	"github.com/bigpandaboy2/project-management-service/internal/app/routes"
)

func main() {
    config.LoadConfig()
    config.ConnectDB()

    router := routes.SetupRouter()

    log.Println("Server is running on port 8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}