package postgres

import (
	"context"
	"database/sql"

	"github.com/namburisnehitha/IssueTracker/domain"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type PostgresIssueRepository struct {
	db     *sql.DB
	tracer trace.Tracer
}

func NewPostgresIssueRepository(db *sql.DB) *PostgresIssueRepository {
	return &PostgresIssueRepository{
		db:     db,
		tracer: otel.Tracer("postgres-issue-repo"),
	}
}

func (ir *PostgresIssueRepository) Save(ctx context.Context, issue domain.Issue) error {

	query := "INSERT into issues (id,title,issue_description,issue_status,created_at,assignee_id) values($1,$2,$3,$4,$5,$6)"

	ctx, span := ir.tracer.Start(ctx, "CreateIssue")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := ir.db.Exec(query, issue.Id, issue.Title, issue.Description, issue.Status, issue.CreatedAt, issue.AssigneeId)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return err
}

func (ir *PostgresIssueRepository) GetById(ctx context.Context, id string) (domain.Issue, error) {

	var issue domain.Issue
	query := `SELECT id,title,issue_description,issue_status,created_at,assignee_id FROM issues WHERE id = $1 `

	ctx, span := ir.tracer.Start(ctx, "GetById")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	err := ir.db.QueryRow(query, id).Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.CreatedAt, &issue.AssigneeId)
	if err != nil {
		span.RecordError(err)
		return domain.Issue{}, err
	}
	return issue, err
}

func (ir *PostgresIssueRepository) GetByTitle(ctx context.Context, title string) ([]domain.Issue, error) {

	var issues []domain.Issue
	query := `SELECT id,title,issue_description,issue_status,created_at,assignee_id FROM issues WHERE title = $1 `

	ctx, span := ir.tracer.Start(ctx, "GetByTitle")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ir.db.Query(query, title)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i domain.Issue
		err := rows.Scan(&i.Id, &i.Title, &i.Description, &i.Status, &i.CreatedAt, &i.AssigneeId)
		if err != nil {
			span.RecordError(err)
			return nil, err
		}
		issues = append(issues, i)
	}
	return issues, err
}

func (ir *PostgresIssueRepository) GetByStatus(ctx context.Context, status domain.IssueStatus) ([]domain.Issue, error) {

	var issues []domain.Issue
	query := `SELECT id,title,issue_description,issue_status,created_at,assignee_id FROM issues WHERE issue_status = $1 `

	ctx, span := ir.tracer.Start(ctx, "GetByStatus")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ir.db.Query(query, status)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i domain.Issue
		err := rows.Scan(&i.Id, &i.Title, &i.Description, &i.Status, &i.CreatedAt, &i.AssigneeId)
		if err != nil {
			span.RecordError(err)
			return nil, err
		}
		issues = append(issues, i)
	}
	return issues, err
}

func (ir *PostgresIssueRepository) UpdateIssue(ctx context.Context, issue domain.Issue) error {

	query := `UPDATE issues SET title =$1,issue_description = $2,issue_status = $3,assignee_id = $4 WHERE id = $5 `

	ctx, span := ir.tracer.Start(ctx, "UpdateIssue")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := ir.db.Exec(query, issue.Title, issue.Description, issue.Status, issue.AssigneeId, issue.Id)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return err
}

func (ir *PostgresIssueRepository) DeleteIssue(ctx context.Context, issue domain.Issue) error {

	query := `DELETE FROM issues where id = $1`

	ctx, span := ir.tracer.Start(ctx, "DeleteIssue")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := ir.db.Exec(query, issue.Id)
	if err != nil {
		span.RecordError(err)
	}
	return err
}

func (ir *PostgresIssueRepository) ListIssues(ctx context.Context) ([]domain.Issue, error) {

	var issues []domain.Issue
	query := `SELECT id,title,issue_description,issue_status,created_at,assignee_id FROM issues `

	ctx, span := ir.tracer.Start(ctx, "ListIssues")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := ir.db.Query(query)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i domain.Issue
		err := rows.Scan(&i.Id, &i.Title, &i.Description, &i.Status, &i.CreatedAt, &i.AssigneeId)
		if err != nil {
			span.RecordError(err)
			return nil, err
		}
		issues = append(issues, i)
	}
	return issues, err
}
