package models

import (
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	DueDatetime time.Time `json:"due_datetime"`
	Contact     string    `json:"contact"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func (t *Task) Bind(r *http.Request) error {
	if t.Title == "" {
		return fmt.Errorf("title is a required field")
	}
	return nil
}

func (*TaskList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Task) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
