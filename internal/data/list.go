package data

import (
	"database/sql"
	"errors"

	"github.com/EmotionlessDev/todo-tasks/internal/validator"
)

type List struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ListModel struct {
	DB *sql.DB
}

func ValidateList(v *validator.Validator, list *List) {
	v.Check(list.Title != "", "title", "must be provided")
	v.Check(len(list.Title) <= 500, "title", "must not be more than 500 bytes long")
}

func (m ListModel) Insert(list *List) error {
	stmt := `
	INSERT INTO list (title)
	VALUES($1)
	RETURNING id
	`

	args := []interface{}{list.Title}
	return m.DB.QueryRow(stmt, args...).Scan(&list.ID)
}

func (m ListModel) Get(id int64) (*List, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	stmt := `
	SELECT id, title
	FROM list
	WHERE id = $1
	`
	list := List{}
	err := m.DB.QueryRow(stmt, id).Scan(
		&list.ID,
		&list.Title,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &list, nil
}

func (m ListModel) Update(id int, title string) error {
	return nil
}

func (m ListModel) Delete(id int) error {
	return nil
}
