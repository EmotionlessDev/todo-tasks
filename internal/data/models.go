package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Task TaskModel
	List ListModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Task: TaskModel{DB: db},
		List: ListModel{DB: db},
	}
}
