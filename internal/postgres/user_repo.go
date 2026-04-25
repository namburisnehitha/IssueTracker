package postgres

import (
	"database/sql"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type PostgresUsersRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUsersRepository {
	return &PostgresUsersRepository{
		db: db,
	}
}

func (ur *PostgresUsersRepository) Save(user domain.User) error {
	query := `INSERT INTO users (id,user_name,user_role,joined_at,changed_role_at) VALUES ($1, $2,$3,$4,$5)`
	_, err := ur.db.Exec(query, user.Id, user.Name, user.Role, user.JoinedAt, user.ChangedRoleAt)
	return err
}
