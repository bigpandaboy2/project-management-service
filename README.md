
# Project Management Service

## Introduction

This project is a RESTful API for a Project Management microservice built in Go. It provides endpoints for managing users, tasks, and projects, including operations such as creating, updating, deleting, and searching.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/bigpandaboy2/project-management-service.git
   cd project-management-service
   ```

2. Build the project:

   ```sh
   make build
   ```

3. Run the server:

   ```sh
   make up
   ```

## API Endpoints

### Users
#### URL: /users

- **GET /users**: Get a list of all users.

- **POST /users**: Create a new user.
  - Request Body:
    ```json
    {
        "full_name": "John Doe",
        "email": "johndoe@example.com",
        "registration": "2024-07-27T00:00:00Z",
        "role": "admin"
    }
    ```

- **GET /users/{id}**: Get details of a specific user.

- **PUT /users/{id}**: Update details of a specific user.
  - Request Body:
    ```json
    {
        "full_name": "John Doe",
        "email": "johndoe@example.com",
        "registration": "2024-07-27T00:00:00Z",
        "role": "manager"
    }
    ```

- **DELETE /users/{id}**: Delete a specific user.

### Projects
#### URL: /projects

- **GET /projects**: Get a list of all projects.

- **POST /projects**: Create a new project.
  - Request Body:
    ```json
    {
        "title": "Project Alpha",
        "description": "A new innovative project",
        "start_date": "2024-07-01T00:00:00Z",
        "end_date": "2024-12-31T00:00:00Z",
        "manager_id": "1"
    }
    ```

- **GET /projects/{id}**: Get details of a specific project.

- **PUT /projects/{id}**: Update details of a specific project.
  - Request Body:
    ```json
    {
        "title": "Project Beta",
        "description": "An updated innovative project",
        "start_date": "2024-07-01T00:00:00Z",
        "end_date": "2024-12-31T00:00:00Z",
        "manager_id": "1"
    }
    ```

- **DELETE /projects/{id}**: Delete a specific project.

### Tasks
#### URL: /tasks

- **GET /tasks**: Get a list of all tasks.

- **POST /tasks**: Create a new task.
  - Request Body:
    ```json
    {
        "title": "Finish Report",
        "description": "Complete the quarterly financial report",
        "priority": "High",
        "state": "InProgress",
        "assignee": "3",
        "project_id": "5",
        "created_at": "2024-07-01T00:00:00Z",
        "completed_at": "2024-07-15T00:00:00Z"
    }
    ```

- **GET /tasks/{id}**: Get details of a specific task.

- **PUT /tasks/{id}**: Update details of a specific task.
  - Request Body:
    ```json
    {
        "title": "Finish Report",
        "description": "Complete the quarterly financial report and review",
        "priority": "High",
        "state": "Completed",
        "assignee": "3",
        "project_id": "5",
        "created_at": "2024-07-01T00:00:00Z",
        "completed_at": "2024-07-10T00:00:00Z"
    }
    ```

- **DELETE /tasks/{id}**: Delete a specific task.