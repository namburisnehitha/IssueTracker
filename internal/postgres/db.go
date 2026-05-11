package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func NewDB(ConnString string) (*sql.DB, error) {

	db, err := otelsql.Open("postgres", ConnString)
	otelsql.WithAttributes(semconv.DBSystemPostgreSQL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}
