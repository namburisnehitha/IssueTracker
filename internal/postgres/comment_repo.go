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

type PostgresCommentRepository struct {
	db     *sql.DB
	tracer trace.Tracer
}

func NewPostgresCommentRepository(db *sql.DB) *PostgresCommentRepository {
	return &PostgresCommentRepository{
		db:     db,
		tracer: otel.Tracer("postgres-comment-repo"),
	}
}

func (cr *PostgresCommentRepository) Save(ctx context.Context, comment domain.Comment) error {

	query := `INSERT into comments ( id,issue_id,user_id,content,created_at,updated_at) values($1,$2,$3,$4,$5,$6)`

	ctx, span := cr.tracer.Start(ctx, "CreateComment")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := cr.db.ExecContext(ctx, query, comment.Id, comment.IssueId, comment.UserId, comment.Content, comment.CreatedAt, comment.UpdatedAt)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return err
	}
	return err
}

func (cr *PostgresCommentRepository) GetById(ctx context.Context, id string) (domain.Comment, error) {

	var comment domain.Comment
	query := `SELECT  id,issue_id,user_id,content,created_at,updated_at FROM comments WHERE id = $1`

	ctx, span := cr.tracer.Start(ctx, "GetById")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	err := cr.db.QueryRowContext(ctx, query, id).Scan(&comment.Id, &comment.IssueId, &comment.UserId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return domain.Comment{}, err
	}
	return comment, err
}

func (cr *PostgresCommentRepository) GetByUserId(ctx context.Context, userid string) ([]domain.Comment, error) {

	var comments []domain.Comment
	query := `SELECT  id,issue_id,user_id,content,created_at,updated_at FROM comments WHERE user_id = $1`

	ctx, span := cr.tracer.Start(ctx, "GetByUserId")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := cr.db.QueryContext(ctx, query, userid)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c domain.Comment
		err = rows.Scan(&c.Id, &c.IssueId, &c.UserId, &c.Content, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			span.SetStatus(codes.Ok, "")
			return nil, err
		}
		comments = append(comments, c)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return nil, err
	}
	return comments, err
}

func (cr *PostgresCommentRepository) GetByIssueId(ctx context.Context, issueid string) ([]domain.Comment, error) {

	var comments []domain.Comment
	query := `SELECT  id,issue_id,user_id,content,created_at,updated_at FROM comments WHERE issue_id = $1`

	ctx, span := cr.tracer.Start(ctx, "GetByIssueId")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := cr.db.QueryContext(ctx, query, issueid)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c domain.Comment
		err = rows.Scan(&c.Id, &c.IssueId, &c.UserId, &c.Content, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			span.SetStatus(codes.Ok, "")
			return nil, err
		}
		comments = append(comments, c)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return nil, err
	}
	return comments, err
}

func (cr *PostgresCommentRepository) UpdateComment(ctx context.Context, comment domain.Comment) error {

	query := `UPDATE comments SET content = $1,updated_at = $2 WHERE id = $3`

	ctx, span := cr.tracer.Start(ctx, "UpdateComment")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := cr.db.ExecContext(ctx, query, comment.Content, comment.UpdatedAt, comment.Id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return err
	}
	return err
}

func (cr *PostgresCommentRepository) DeleteComment(ctx context.Context, comment domain.Comment) error {

	query := `DELETE FROM comments WHERE id = $1`

	ctx, span := cr.tracer.Start(ctx, "DeleteComment")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := cr.db.ExecContext(ctx, query, comment.Id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return err
	}
	return err
}

func (cr *PostgresCommentRepository) CommentList(ctx context.Context) ([]domain.Comment, error) {

	var comments []domain.Comment
	query := `SELECT  id,issue_id,user_id,content,created_at,updated_at FROM comments `

	ctx, span := cr.tracer.Start(ctx, "CommentList")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c domain.Comment
		err = rows.Scan(&c.Id, &c.IssueId, &c.UserId, &c.Content, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			span.SetStatus(codes.Ok, "")
			return nil, err
		}
		comments = append(comments, c)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetStatus(codes.Ok, "")
		return nil, err
	}
	return comments, err
}
