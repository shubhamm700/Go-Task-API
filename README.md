# Task API

## Introduction
Task API is a simple API service designed to manage tasks, enabling users to create, retrieve, update, and delete tasks with associated details such as title, description, priority, due date, and contact information.

## Requirements
- Docker
- Go
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

## Getting Started
1. Clone the repository:
    ```bash
    git clone gitlab.com/idoko/task-api
    ```

2. Copy the `env.example` file to a `.env` file:
    ```bash
    cp .env.example .env
    ```

3. Update the PostgreSQL variables in the `.env` file to match your preferences.

4. Build and start the services with Docker:
    ```bash
    docker-compose up --build
    ```

5. Apply database migrations using `migrate`:
    ```bash
    export POSTGRESQL_URL="postgres://$PG_USER:$PG_PASS@localhost:5432/$PG_DB?sslmode=disable"
    migrate -database ${POSTGRESQL_URL} -path db/migrations up
    ```
    _**NOTE:** Replace the `$PG*` variables with actual values._

## Development
After making changes, rebuild the `server` service:
```bash
docker-compose stop server
docker-compose build server
docker-compose up --no-start server
docker-compose start server
```

### Base URL

The base URL for the Task API is `http://localhost:8080` (assuming the API is running locally on port 8080).

### API Endpoints

#### 1. Get All Tasks

- **Endpoint:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieve a list of all tasks.
- **Request:**
  - Headers: None
- **Response:**
  - Status Code: `200 OK`
  - Body: List of tasks in JSON format.

#### 2. Get Task by ID

- **Endpoint:** `/tasks/{id}`
- **Method:** `GET`
- **Description:** Retrieve details of a specific task by its ID.
- **Request:**
  - Headers: None
- **Response:**
  - Status Code: `200 OK`
  - Body: Task details in JSON format.
  - Status Code: `404 Not Found` if the task with the specified ID does not exist.

#### 3. Create a New Task

- **Endpoint:** `/tasks`
- **Method:** `POST`
- **Description:** Create a new task.
- **Request:**
  - Headers:
    - `Content-Type: application/json`
  - Body: JSON object representing the new task.
- **Response:**
  - Status Code: `201 Created`
  - Body: Details of the newly created task in JSON format.

#### 4. Update Task by ID

- **Endpoint:** `/tasks/{id}`
- **Method:** `PUT`
- **Description:** Update details of a specific task by its ID.
- **Request:**
  - Headers:
    - `Content-Type: application/json`
  - Body: JSON object with updated task details.
- **Response:**
  - Status Code: `200 OK`
  - Body: Details of the updated task in JSON format.
  - Status Code: `404 Not Found` if the task with the specified ID does not exist.

#### 5. Delete Task by ID

- **Endpoint:** `/tasks/{id}`
- **Method:** `DELETE`
- **Description:** Delete a specific task by its ID.
- **Request:**
  - Headers: None
- **Response:**
  - Status Code: `204 No Content`
  - Status Code: `404 Not Found` if the task with the specified ID does not exist.
