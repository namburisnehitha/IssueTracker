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
	query := "INSERT into issues (id,title,issue_description,issue_status,created_at,assignee_id) values($1,$2,$3,$4,$5,$6)"
	_, err := ir.db.Exec(query, issue.Id, issue.Title, issue.Description, issue.Status, issue.CreatedAt, issue.AssigneeId)
	return err
}

func (ir *PostgresIssueRepository) GetById(id string) (domain.Issue, error) {
	var issue domain.Issue
	query := `SELECT id,title,issue_description,issue_status,created_at,assignee_id FROM issues WHERE id = $1 `
	err := ir.db.QueryRow(query, id).Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Status, &issue.CreatedAt, &issue.AssigneeId)
	return issue, err
}

func (ir *PostgresIssueRepository) GetByTitle(title string) ([]domain.Issue, error) {
	var issues []domain.Issue
	query := `SELECT id,title,issue_description,issue_status,created_at,assignee_id FROM issues WHERE title = $1 `
	rows, err := ir.db.Query(query, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i domain.Issue
		err := rows.Scan(&i.Id, &i.Title, &i.Description, &i.Status, &i.CreatedAt, &i.AssigneeId)
		if err != nil {
			return nil, err
		}
		issues = append(issues, i)
	}
	return issues, err
}

func (ir *PostgresIssueRepository) GetByStatus(status domain.IssueStatus) ([]domain.Issue, error) {
	var issues []domain.Issue
	query := `SELECT id,title,issue_description,issue_status,created_at,assignee_id FROM issues WHERE issue_status = $1 `
	rows, err := ir.db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i domain.Issue
		err := rows.Scan(&i.Id, &i.Title, &i.Description, &i.Status, &i.CreatedAt, &i.AssigneeId)
		if err != nil {
			return nil, err
		}
		issues = append(issues, i)
	}
	return issues, err
}

func (ir *PostgresIssueRepository) UpdateIssue(issue domain.Issue) error {
	query := `UPDATE issues SET title =$1,issue_description = $2,issue_status = $3,assignee_id = $4 WHERE id = $5 `
	_, err := ir.db.Exec(query, issue.Title, issue.Description, issue.Status, issue.AssigneeId, issue.Id)
	return err
}

func (ir *PostgresIssueRepository) DeleteIssue(issue domain.Issue) error {
	query := `DELETE FROM issues where id = $1`
	_, err := ir.db.Exec(query, issue.Id)
	return err
}

func (ir *PostgresIssueRepository) ListIssues() ([]domain.Issue, error) {
	var issues []domain.Issue
	query := `SELECT id,title,issue_description,issue_status,created_at,assignee_id FROM issues `
	rows, err := ir.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var i domain.Issue
		err := rows.Scan(&i.Id, &i.Title, &i.Description, &i.Status, &i.CreatedAt, &i.AssigneeId)
		if err != nil {
			return nil, err
		}
		issues = append(issues, i)
	}
	return issues, err
}
