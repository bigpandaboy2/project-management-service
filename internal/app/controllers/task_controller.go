package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/bigpandaboy2/project-management-service/internal/app/models"
    "github.com/bigpandaboy2/project-management-service/internal/app/services"
    "github.com/bigpandaboy2/project-management-service/internal/app/utils"
    "github.com/google/uuid"
    "github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := services.GetAllTasks()
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }
    utils.JSONResponse(w, tasks, http.StatusOK)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    task.ID = uuid.New()
    if err := services.CreateTask(&task); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, task, http.StatusCreated)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    task, err := services.GetTask(id)
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusNotFound)
        return
    }

    utils.JSONResponse(w, task, http.StatusOK)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    task.ID = id
    if err := services.UpdateTask(&task); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, task, http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.DeleteTask(id); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, map[string]string{"message": "Task deleted successfully"}, http.StatusOK)
}