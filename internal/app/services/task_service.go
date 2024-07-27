package services

import (
	"database/sql"
    
	"github.com/bigpandaboy2/project-management-service/internal/app/config"
	"github.com/bigpandaboy2/project-management-service/internal/app/models"
	"github.com/google/uuid"
)

func GetAllTasks() ([]models.Task, error) {
    rows, err := config.DB.Query("SELECT id, title, description, priority, state, assignee, project_id, created_at, completed_at FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.State, &task.Assignee, &task.ProjectID, &task.CreatedAt, &task.CompletedAt); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func CreateTask(task *models.Task) error {
    _, err := config.DB.Exec("INSERT INTO tasks (id, title, description, priority, state, assignee, project_id, created_at, completed_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
        task.ID, task.Title, task.Description, task.Priority, task.State, task.Assignee, task.ProjectID, task.CreatedAt, task.CompletedAt)
    return err
}

func GetTask(id uuid.UUID) (*models.Task, error) {
    var task models.Task
    err := config.DB.QueryRow("SELECT id, title, description, priority, state, assignee, project_id, created_at, completed_at FROM tasks WHERE id = $1", id).
        Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.State, &task.Assignee, &task.ProjectID, &task.CreatedAt, &task.CompletedAt)
    if err == sql.ErrNoRows {
        return nil, err
    }
    return &task, err
}

func UpdateTask(task *models.Task) error {
    _, err := config.DB.Exec("UPDATE tasks SET title = $1, description = $2, priority = $3, state = $4, assignee = $5, project_id = $6, created_at = $7, completed_at = $8 WHERE id = $9",
        task.Title, task.Description, task.Priority, task.State, task.Assignee, task.ProjectID, task.CreatedAt, task.CompletedAt, task.ID)
    return err
}

func DeleteTask(id uuid.UUID) error {
    _, err := config.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
    return err
}