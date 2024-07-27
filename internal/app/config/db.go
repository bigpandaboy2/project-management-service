package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/lib/pq"
    "github.com/cenkalti/backoff/v4"
)

var DB *sql.DB

func ConnectDB() {
    var err error

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    operation := func() error {
        DB, err = sql.Open("postgres", connStr)
        if err != nil {
            return fmt.Errorf("error opening database: %w", err)
        }

        if err = DB.Ping(); err != nil {
            DB.Close()
            return fmt.Errorf("error connecting to the database: %w", err)
        }

        return nil
    }

    notify := func(err error, duration time.Duration) {
        log.Printf("Error connecting to database: %v. Retrying in %s", err, duration)
    }

    backOff := backoff.NewExponentialBackOff()
    backOff.MaxElapsedTime = 5 * time.Minute

    err = backoff.RetryNotify(operation, backOff, notify)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    log.Println("Successfully connected to the database")
}