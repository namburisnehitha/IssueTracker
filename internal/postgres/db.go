package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewDB(ConnString string) (*sql.DB, error) {

	db, err := sql.Open("postgres", ConnString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}
