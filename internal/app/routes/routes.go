package routes

import (
	"github.com/bigpandaboy2/project-management-service/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
    userGroup := router.Group("/users")
    {
        userGroup.GET("/", controllers.GetUsers)
        userGroup.POST("/", controllers.CreateUser)
        userGroup.GET("/:id", controllers.GetUser)
        userGroup.PUT("/:id", controllers.UpdateUser)
        userGroup.DELETE("/:id", controllers.DeleteUser)
    }

    taskGroup := router.Group("/tasks")
    {
        taskGroup.GET("/", controllers.GetTasks)
        taskGroup.POST("/", controllers.CreateTask)
        taskGroup.GET("/:id", controllers.GetTask)
        taskGroup.PUT("/:id", controllers.UpdateTask)
        taskGroup.DELETE("/:id", controllers.DeleteTask)
    }

    projectGroup := router.Group("/projects")
    {
        projectGroup.GET("/", controllers.GetProjects)
        projectGroup.POST("/", controllers.CreateProject)
        projectGroup.GET("/:id", controllers.GetProject)
        projectGroup.PUT("/:id", controllers.UpdateProject)
        projectGroup.DELETE("/:id", controllers.DeleteProject)
    }
}