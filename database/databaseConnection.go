package database

import (
	"database/sql"
)

func connection() (*sql.DB, error) {
	stringConnectionDataBase := "golang:golang@/devbook?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", stringConnectionDataBase)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err

	}
		return db, nil
}
