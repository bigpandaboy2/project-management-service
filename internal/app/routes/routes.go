package routes

import (
    "github.com/gorilla/mux"
    "github.com/bigpandaboy2/project-management-service/internal/app/controllers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
    r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
    
    r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
    r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
    r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
    
    r.HandleFunc("/projects", controllers.GetProjects).Methods("GET")
    r.HandleFunc("/projects", controllers.CreateProject).Methods("POST")
    r.HandleFunc("/projects/{id}", controllers.GetProject).Methods("GET")
    r.HandleFunc("/projects/{id}", controllers.UpdateProject).Methods("PUT")
    r.HandleFunc("/projects/{id}", controllers.DeleteProject).Methods("DELETE")

    return r
}