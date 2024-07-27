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

func GetProjects(w http.ResponseWriter, r *http.Request) {
    projects, err := services.GetAllProjects()
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }
    utils.JSONResponse(w, projects, http.StatusOK)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
    var project models.Project
    if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    project.ID = uuid.New()
    if err := services.CreateProject(&project); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, project, http.StatusCreated)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    project, err := services.GetProject(id)
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusNotFound)
        return
    }

    utils.JSONResponse(w, project, http.StatusOK)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    var project models.Project
    if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    project.ID = id
    if err := services.UpdateProject(&project); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, project, http.StatusOK)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.DeleteProject(id); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, map[string]string{"message": "Project deleted successfully"}, http.StatusOK)
}