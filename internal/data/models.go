package data

import "database/sql"

type Models struct {
	Task TaskModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Task: TaskModel{DB: db},
	}
}
