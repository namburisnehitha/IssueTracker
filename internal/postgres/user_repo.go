package postgres

import (
	"context"
	"database/sql"

	"github.com/namburisnehitha/IssueTracker/domain"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type PostgresUsersRepository struct {
	db     *sql.DB
	tracer trace.Tracer
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUsersRepository {
	return &PostgresUsersRepository{
		db:     db,
		tracer: otel.Tracer("postgres-user-repo"),
	}
}

func (ur *PostgresUsersRepository) Save(ctx context.Context, user domain.User) error {

	query := `INSERT INTO users (id,user_name,user_role,joined_at,changed_role_at,user_username,user_password) VALUES ($1, $2,$3,$4,$5,$6,$7)`

	ctx, span := ur.tracer.Start(ctx, "CreateUser")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := ur.db.ExecContext(ctx, query, user.Id, user.Name, user.Role, user.JoinedAt, user.ChangedRoleAt, user.UserName, user.Password)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return err
}

func (ur *PostgresUsersRepository) GetById(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	query := `SELECT id,user_name,user_role,joined_at,changed_role_at,user_username,user_password FROM users WHERE id = $1`

	ctx, span := ur.tracer.Start(ctx, "GetById")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	err := ur.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Name, &user.Role, &user.JoinedAt, &user.ChangedRoleAt, &user.UserName, &user.Password)
	if err != nil {
		span.RecordError(err)
		return domain.User{}, err
	}
	return user, err
}

func (ur *PostgresUsersRepository) GetByName(ctx context.Context, name string) ([]domain.User, error) {
	var users []domain.User
	query := `SELECT id,user_name,user_role,joined_at,changed_role_at,user_username,user_password  FROM users WHERE user_name =$1`

	ctx, span := ur.tracer.Start(ctx, "GetByName")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ur.db.QueryContext(ctx, query, name)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u domain.User
		err = rows.Scan(&u.Id, &u.Name, &u.Role, &u.JoinedAt, &u.ChangedRoleAt, &u.UserName, &u.Password)
		if err != nil {
			span.RecordError(err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, err
}

func (ur *PostgresUsersRepository) GetByRole(ctx context.Context, role domain.Roles) ([]domain.User, error) {

	var users []domain.User
	query := `SELECT id,user_name,user_role,joined_at,changed_role_at,user_username,user_password  FROM users WHERE user_role =$1`

	ctx, span := ur.tracer.Start(ctx, "GetByRole")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ur.db.QueryContext(ctx, query, role)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u domain.User
		err = rows.Scan(&u.Id, &u.Name, &u.Role, &u.JoinedAt, &u.ChangedRoleAt, &u.UserName, &u.Password)
		if err != nil {
			span.RecordError(err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, err
}

func (ur *PostgresUsersRepository) UpdateUser(ctx context.Context, user domain.User) error {

	query := `UPDATE users SET user_name = $1, user_role = $2,changed_role_at = $3 WHERE id = $4 `

	ctx, span := ur.tracer.Start(ctx, "UpdateUser")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := ur.db.ExecContext(ctx, query, user.Name, user.Role, user.ChangedRoleAt, user.Id)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return err
}

func (ur *PostgresUsersRepository) DeleteUser(ctx context.Context, user domain.User) error {
	query := `DELETE FROM users WHERE id = $1 `

	ctx, span := ur.tracer.Start(ctx, "DeleteUser")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := ur.db.ExecContext(ctx, query, user.Id)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return err
}

func (ur *PostgresUsersRepository) UserList(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	query := `SELECT id,user_name,user_role,joined_at,changed_role_at,user_username,user_password FROM users `

	ctx, span := ur.tracer.Start(ctx, "UserList")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u domain.User
		err = rows.Scan(&u.Id, &u.Name, &u.Role, &u.JoinedAt, &u.ChangedRoleAt, &u.UserName, &u.Password)
		if err != nil {
			span.RecordError(err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, err
}

func (ur *PostgresUsersRepository) GetByUserName(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	query := `SELECT id,user_name,user_role,joined_at,changed_role_at,user_username,user_password FROM users WHERE user_username = $1`

	ctx, span := ur.tracer.Start(ctx, "GetByUserName")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	err := ur.db.QueryRowContext(ctx, query, username).Scan(&user.Id, &user.Name, &user.Role, &user.JoinedAt, &user.ChangedRoleAt, &user.UserName, &user.Password)
	if err != nil {
		span.RecordError(err)
		return domain.User{}, err
	}
	return user, err
}
