package postgres

import (
	"database/sql"
)

func NewDB(ConnString string) (*sql.DB, error) {

	db, err := sql.Open("postgres", ConnString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}
