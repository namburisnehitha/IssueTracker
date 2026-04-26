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

func (ur *PostgresUsersRepository) GetById(id string) (domain.User, error) {
	var user domain.User
	query := `SELECT id,user_name,user_role,joined_at,changed_role_at FROM users WHERE id = $1`
	err := ur.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Role, &user.JoinedAt, &user.ChangedRoleAt)

	return user, err
}

func (ur *PostgresUsersRepository) GetByName(name string) ([]domain.User, error) {
	var users []domain.User
	query := `SELECT id,user_name,user_role,joined_at,changed_role_at FROM users WHERE user_name =$1`
	rows, err := ur.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u domain.User
		err = rows.Scan(&u.Id, &u.Name, &u.Role, &u.JoinedAt, &u.ChangedRoleAt)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, err
}
