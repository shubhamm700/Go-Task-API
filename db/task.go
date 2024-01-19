package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/shubhamm700/Go-Task-API/models"
)
func (db Database) GetAllTasks() (*models.TaskList, error) {
	list := &models.TaskList{}
	rows, err := db.Conn.Query("SELECT * FROM tasks ORDER BY id DESC")
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.DueDatetime, &task.Contact, &task.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Tasks = append(list.Tasks, task)
	}
	return list, nil
}

// GetTaskById retrieves a task from the database by its ID
func (db Database) GetTaskById(taskId int) (models.Task, error) {
	task := models.Task{}
	query := `SELECT * FROM tasks WHERE id = $1;`
	row := db.Conn.QueryRow(query, taskId)
	switch err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.DueDatetime, &task.Contact, &task.CreatedAt); err {
	case sql.ErrNoRows:
		return task, ErrNoMatch
	default:
		return task, err
	}
}

// AddTask adds a new task to the database
func (db Database) AddTask(task *models.Task) error {
	var id int
	var createdAt time.Time

	query := `
		INSERT INTO tasks (title, description, priority, due_datetime, contact)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	err := db.Conn.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Priority,
		task.DueDatetime,
		task.Contact,
	).Scan(&id, &createdAt)

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			// Duplicate key violation (unique constraint)
			return fmt.Errorf("a task with the same title already exists")
		}
		return err
	}

	task.ID = id
	task.CreatedAt = createdAt
	return nil
}

// UpdateTaskById updates a task in the database by its ID
func (db Database) UpdateTaskById(taskId int, taskData models.Task) (models.Task, error) {
	task := models.Task{}
	query := `
		UPDATE tasks
		SET title=$1, description=$2, priority=$3, due_datetime=$4, contact=$5
		WHERE id=$6
		RETURNING id, title, description, priority, due_datetime, contact, created_at;
	`

	err := db.Conn.QueryRow(
		query,
		taskData.Title,
		taskData.Description,
		taskData.Priority,
		taskData.DueDatetime,
		taskData.Contact,
		taskId,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.DueDatetime, &task.Contact, &task.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return task, ErrNoMatch
		}
		return task, err
	}

	return task, nil
}

// DeleteTaskById deletes a task from the database by its ID
func (db Database) DeleteTaskById(taskId int) error {
	query := `DELETE FROM tasks WHERE id = $1;`
	result, err := db.Conn.Exec(query, taskId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNoMatch
	}

	return nil
}