package data

import (
	"database/sql"
	"time"

	"github.com/EmotionlessDev/todo-tasks/internal/validator"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Completed   bool      `json:"completed"`
}

type TaskModel struct {
	DB *sql.DB
}

func ValidateTask(v *validator.Validator, task *Task) {
	v.Check(task.Title != "", "title", "must be provided")
	v.Check(len(task.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(task.Description != "", "description", "must be provided")
}

func (m TaskModel) Insert(title, description string) (int, error) {
	return 0, nil
}

func (m TaskModel) Get(id int) (*Task, error) {
	return nil, nil
}

func (m TaskModel) Update(id int, title, description string) error {
	return nil
}

func (m TaskModel) Delete(id int) error {
	return nil
}
