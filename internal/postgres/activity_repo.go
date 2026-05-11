package postgres

import (
	"context"
	"database/sql"

	"github.com/namburisnehitha/IssueTracker/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type PostgresActivityRepository struct {
	db     *sql.DB
	tracer trace.Tracer
}

func NewPostgresActivityRepository(db *sql.DB) *PostgresActivityRepository {
	return &PostgresActivityRepository{
		db:     db,
		tracer: otel.Tracer("postgres-activity-repo"),
	}
}

func (ar *PostgresActivityRepository) Save(ctx context.Context, activity domain.Activity) error {

	query := `INSERT into activities(id,issue_id,user_id,activity_description,created_at,activity_action) values($1,$2,$3,$4,$5,$6)`

	ctx, span := ar.tracer.Start(ctx, "Createactivity")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := ar.db.ExecContext(ctx, query, activity.Id, activity.IssueId, activity.UserId, activity.Description, activity.CreatedAt, activity.Action)

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return err
}

func (ar *PostgresActivityRepository) GetById(ctx context.Context, id string) (domain.Activity, error) {

	var activity domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities WHERE id = $1`

	ctx, span := ar.tracer.Start(ctx, "GetById")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	err := ar.db.QueryRowContext(ctx, query, id).Scan(&activity.Id, &activity.IssueId, &activity.UserId, &activity.Description, &activity.CreatedAt, &activity.Action)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return domain.Activity{}, err
	}
	return activity, err
}

func (ar *PostgresActivityRepository) GetByUserId(ctx context.Context, userid string) ([]domain.Activity, error) {

	var activities []domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities WHERE user_id = $1`

	ctx, span := ar.tracer.Start(ctx, "GetByUserId")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ar.db.QueryContext(ctx, query, userid)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a domain.Activity
		err := rows.Scan(&a.Id, &a.IssueId, &a.UserId, &a.Description, &a.CreatedAt, &a.Action)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		activities = append(activities, a)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return activities, err
}

func (ar *PostgresActivityRepository) GetByIssueId(ctx context.Context, issueid string) ([]domain.Activity, error) {

	var activities []domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities WHERE issue_id = $1`

	ctx, span := ar.tracer.Start(ctx, "GetByIssueId")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ar.db.QueryContext(ctx, query, issueid)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a domain.Activity
		err := rows.Scan(&a.Id, &a.IssueId, &a.UserId, &a.Description, &a.CreatedAt, &a.Action)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		activities = append(activities, a)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return activities, err
}

func (ar *PostgresActivityRepository) GetByAction(ctx context.Context, action domain.ActivityType) ([]domain.Activity, error) {

	var activities []domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities WHERE activity_action = $1`

	ctx, span := ar.tracer.Start(ctx, "GetByAction")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ar.db.QueryContext(ctx, query, action)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	for rows.Next() {
		var a domain.Activity
		err := rows.Scan(&a.Id, &a.IssueId, &a.UserId, &a.Description, &a.CreatedAt, &a.Action)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		activities = append(activities, a)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return activities, err
}
func (ar *PostgresActivityRepository) ActivityList(ctx context.Context) ([]domain.Activity, error) {

	var activities []domain.Activity
	query := `SELECT id,issue_id,user_id,activity_description,created_at,activity_action FROM activities`

	ctx, span := ar.tracer.Start(ctx, "ActivityList")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ar.db.QueryContext(ctx, query)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	for rows.Next() {
		var a domain.Activity
		err := rows.Scan(&a.Id, &a.IssueId, &a.UserId, &a.Description, &a.CreatedAt, &a.Action)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		activities = append(activities, a)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return activities, err
}
