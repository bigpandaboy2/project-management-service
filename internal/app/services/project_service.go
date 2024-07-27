package services

import (
	"database/sql"

	"github.com/bigpandaboy2/project-management-service/internal/app/config"
	"github.com/bigpandaboy2/project-management-service/internal/app/models"
	"github.com/google/uuid"
)

func GetAllProjects() ([]models.Project, error) {
    rows, err := config.DB.Query("SELECT id, title, description, start_date, end_date, manager_id FROM projects")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []models.Project
    for rows.Next() {
        var project models.Project
        if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID); err != nil {
            return nil, err
        }
        projects = append(projects, project)
    }

    return projects, nil
}

func CreateProject(project *models.Project) error {
    _, err := config.DB.Exec("INSERT INTO projects (id, title, description, start_date, end_date, manager_id) VALUES ($1, $2, $3, $4, $5, $6)",
        project.ID, project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID)
    return err
}

func GetProject(id uuid.UUID) (*models.Project, error) {
    var project models.Project
    err := config.DB.QueryRow("SELECT id, title, description, start_date, end_date, manager_id FROM projects WHERE id = $1", id).
        Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID)
    if err == sql.ErrNoRows {
        return nil, err
    }
    return &project, err
}

func UpdateProject(project *models.Project) error {
    _, err := config.DB.Exec("UPDATE projects SET title = $1, description = $2, start_date = $3, end_date = $4, manager_id = $5 WHERE id = $6",
        project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID, project.ID)
    return err
}

func DeleteProject(id uuid.UUID) error {
    _, err := config.DB.Exec("DELETE FROM projects WHERE id = $1", id)
    return err
}