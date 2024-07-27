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

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := services.GetAllUsers()
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }
    utils.JSONResponse(w, users, http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.ID = uuid.New()
    if err := services.CreateUser(&user); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, user, http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := services.GetUser(id)
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusNotFound)
        return
    }

    utils.JSONResponse(w, user, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.ID = id
    if err := services.UpdateUser(&user); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, user, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := uuid.Parse(params["id"])
    if err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.DeleteUser(id); err != nil {
        utils.JSONResponse(w, err.Error(), http.StatusInternalServerError)
        return
    }

    utils.JSONResponse(w, map[string]string{"message": "User deleted successfully"}, http.StatusOK)
}