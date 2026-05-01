package postgres

import (
	"database/sql"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type PostgresCommentRepository struct {
	db *sql.DB
}

func NewPostgresCommentRepository(db *sql.DB) *PostgresCommentRepository {
	return &PostgresCommentRepository{
		db: db,
	}
}

func (cr *PostgresCommentRepository) Save(comment domain.Comment) error {
	query := `INSERT into comments ( id,issue_id,user_id,content,created_at,updated_at) values($1,$2,$3,$4,$5,$6)`
	_, err := cr.db.Exec(query, comment.Id, comment.IssueId, comment.UserId, comment.Content, comment.CreatedAt, comment.UpdatedAt)
	return err
}

func (cr *PostgresCommentRepository) GetById(id string) (domain.Comment, error) {
	var comment domain.Comment
	query := `SELECT  id,issue_id,user_id,content,created_at,updated_at FROM comments WHERE id = $1`
	err := cr.db.QueryRow(query, id).Scan(&comment.Id, &comment.IssueId, &comment.UserId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	return comment, err
}

func (cr *PostgresCommentRepository) GetByUserId(userid string) ([]domain.Comment, error) {
	var comments []domain.Comment
	query := `SELECT  id,issue_id,user_id,content,created_at,updated_at FROM comments WHERE user_id = $1`
	rows, err := cr.db.Query(query, userid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c domain.Comment
		err = rows.Scan(&c.Id, &c.IssueId, &c.UserId, &c.Content, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, err
}

func (cr *PostgresCommentRepository) GetByIssueId(issueid string) ([]domain.Comment, error) {
	var comments []domain.Comment
	query := `SELECT  id,issue_id,user_id,content,created_at,updated_at FROM comments WHERE issue_id = $1`
	rows, err := cr.db.Query(query, issueid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c domain.Comment
		err = rows.Scan(&c.Id, &c.IssueId, &c.UserId, &c.Content, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, err
}

func (cr *PostgresCommentRepository) UpdateComment(comment domain.Comment) error {
	query := `UPDATE comments SET content = $1,updated_at = $2 WHERE id = $3`
	_, err := cr.db.Exec(query, comment.Content, comment.UpdatedAt, comment.Id)
	return err
}

func (cr *PostgresCommentRepository) DeleteComment(comment domain.Comment) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := cr.db.Exec(query, comment.Id)
	return err
}

func (cr *PostgresCommentRepository) CommentList() ([]domain.Comment, error) {
	var comments []domain.Comment
	query := `SELECT  id,issue_id,user_id,content,created_at,updated_at FROM comments `
	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c domain.Comment
		err = rows.Scan(&c.Id, &c.IssueId, &c.UserId, &c.Content, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, err
}
