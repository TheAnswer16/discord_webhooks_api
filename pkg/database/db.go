package database

import (
	"database/sql"
	_"github.com/lib/pq"
)

func NewConnection(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}