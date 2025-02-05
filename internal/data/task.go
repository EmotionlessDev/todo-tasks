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
	ListID      int       `json:"list_id"`
}

type TaskModel struct {
	DB *sql.DB
}

func ValidateTask(v *validator.Validator, task *Task) {
	v.Check(task.Title != "", "title", "must be provided")
	v.Check(len(task.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(task.Description != "", "description", "must be provided")
	v.Check(len(task.Description) <= 1000, "description", "must not be more than 1000 bytes long")
	v.Check(task.ListID > 0, "list_id", "must be a positive integer")
}

func (m TaskModel) Insert(title, description string, listID int) error {
	stmt := `
	INSERT INTO task (title, description, created_at, completed, list_id)
	VALUES($1, $2, $3, $4, $5)
	`
	task := &Task{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
		ListID:      listID,
	}

	_, err := m.DB.Exec(stmt, task.Title, task.Description, task.CreatedAt, task.Completed, task.ListID)
	if err != nil {
		return err
	}
	return nil
}

func (m TaskModel) Get(id int64) (*Task, error) {
	stmt := `
	SELECT id, title, description, created_at, completed
	FROM task
	WHERE id = $1
	`
	task := &Task{}
	err := m.DB.QueryRow(stmt, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.CreatedAt,
		&task.Completed,
	)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (m TaskModel) Update(id int, title, description string) error {
	return nil
}

func (m TaskModel) Delete(id int) error {
	return nil
}
