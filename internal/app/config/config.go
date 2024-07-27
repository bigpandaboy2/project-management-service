package config

import (
    "log"
    "github.com/joho/godotenv"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using default settings")
    } else {
        log.Println("Configuration loaded")
    }
}