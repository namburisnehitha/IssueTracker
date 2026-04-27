package postgres

import (
	"database/sql"
	"github.com/namburisnehitha/IssueTracker/domain"
)

type PostgresIssueRepository struct {
	db *sql.DB
}

func NewPostgresIssueRepository(db *sql.DB) *PostgresIssueRepository {
	return &PostgresIssueRepository{
		db: db,
	}
}

func (ir *PostgresIssueRepository) Save(issue domain.Issue) error {
	query := "INSERT into issues (id,title,issue_description,issue_status,created_at,assignee_id) values($1,$2,$3,$4,$5)"
	_, err := ir.db.Exec(query, issue.Id, issue.Title, issue.Description, issue.Status, issue.CreatedAt, issue.AssigneeId)
	return err
}
