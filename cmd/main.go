package main

import (
    "log"
    "github.com/bigpandaboy2/project-management-service/internal/app/config"
    "github.com/bigpandaboy2/project-management-service/internal/app/routes"
    _ "github.com/bigpandaboy2/project-management-service/docs"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func main() {
    config.LoadConfig()
    config.ConnectDB()

    router := gin.Default()
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    routes.SetupRouter(router)

    log.Println("Server is running on port 8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}