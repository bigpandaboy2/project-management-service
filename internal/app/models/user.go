package models

import (
    "github.com/google/uuid"
    "time"
)

type User struct {
    ID           uuid.UUID `json:"id"`
    FullName     string    `json:"full_name"`
    Email        string    `json:"email"`
    Registration time.Time `json:"registration"`
    Role         string    `json:"role"`
}