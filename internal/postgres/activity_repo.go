package postgres

import (
	"context"
	"database/sql"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type PostgresActivityRepository struct {
	db *sql.DB
}

func NewPostgresActivityRepository(db *sql.DB) *PostgresActivityRepository {
	return &PostgresActivityRepository{
		db: db,
	}
}

func (ar *PostgresActivityRepository) Save(ctx context.Context, activity domain.Activity) error {
	query := `INSERT into activities(id,issue_id,user_id,activity_description,created_at,activity_action) values($1,$2,$3,$4,$5,$6)`
	_, err := ar.db.Exec(query, activity.Id, activity.IssueId, activity.UserId, activity.Description, activity.CreatedAt, activity.Action)
	return err
}

func (ar *PostgresActivityRepository) GetById(ctx context.Context, id string) (domain.Activity, error) {
	var activity domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities WHERE id = $1`
	err := ar.db.QueryRow(query, id).Scan(&activity.Id, &activity.IssueId, &activity.UserId, &activity.Description, &activity.CreatedAt, &activity.Action)
	return activity, err
}

func (ar *PostgresActivityRepository) GetByUserId(ctx context.Context, userid string) ([]domain.Activity, error) {
	var activities []domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities WHERE user_id = $1`
	rows, err := ar.db.Query(query, userid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a domain.Activity
		err := rows.Scan(&a.Id, &a.IssueId, &a.UserId, &a.Description, &a.CreatedAt, &a.Action)
		if err != nil {
			return nil, err
		}
		activities = append(activities, a)
	}
	return activities, err
}

func (ar *PostgresActivityRepository) GetByIssueId(ctx context.Context, issueid string) ([]domain.Activity, error) {
	var activities []domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities WHERE issue_id = $1`
	rows, err := ar.db.Query(query, issueid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a domain.Activity
		err := rows.Scan(&a.Id, &a.IssueId, &a.UserId, &a.Description, &a.CreatedAt, &a.Action)
		if err != nil {
			return nil, err
		}
		activities = append(activities, a)
	}
	return activities, err
}

func (ar *PostgresActivityRepository) GetByAction(ctx context.Context, action domain.ActivityType) ([]domain.Activity, error) {
	var activities []domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities WHERE activity_action = $1`
	rows, err := ar.db.Query(query, action)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a domain.Activity
		err := rows.Scan(&a.Id, &a.IssueId, &a.UserId, &a.Description, &a.CreatedAt, &a.Action)
		if err != nil {
			return nil, err
		}
		activities = append(activities, a)
	}
	return activities, err
}
func (ar *PostgresActivityRepository) ActivityList(ctx context.Context) ([]domain.Activity, error) {
	var activities []domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities`
	rows, err := ar.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a domain.Activity
		err := rows.Scan(&a.Id, &a.IssueId, &a.UserId, &a.Description, &a.CreatedAt, &a.Action)
		if err != nil {
			return nil, err
		}
		activities = append(activities, a)
	}
	return activities, err
}
