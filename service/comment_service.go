package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/namburisnehitha/IssueTracker/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type CommentService struct {
	commentRepository domain.CommentRepository
	tracer            trace.Tracer
}

func NewCommentService(commentRepository domain.CommentRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
		tracer:            otel.Tracer("comment-service"),
	}
}

func (c *CommentService) CreateComment(ctx context.Context, issueid string, userid string, content string) (string, error) {

	ctx, span := c.tracer.Start(ctx, "CreateComment")
	defer span.End()

	id := uuid.New().String()
	comment, err := domain.NewComment(issueid, userid, content, id)
	if err != nil {
		return "", err
	}
	return id, c.commentRepository.Save(ctx, comment)
}

func (c *CommentService) GetByIssueId(ctx context.Context, issueid string) ([]domain.Comment, error) {
	ctx, span := c.tracer.Start(ctx, "GetByIssueId")
	defer span.End()
	return c.commentRepository.GetByIssueId(ctx, issueid)
}

func (c *CommentService) GetById(ctx context.Context, id string) (domain.Comment, error) {
	ctx, span := c.tracer.Start(ctx, "GetById")
	defer span.End()
	return c.commentRepository.GetById(ctx, id)
}

func (c *CommentService) GetByUserId(ctx context.Context, userid string) ([]domain.Comment, error) {
	ctx, span := c.tracer.Start(ctx, "GetByUserId")
	defer span.End()
	return c.commentRepository.GetByUserId(ctx, userid)
}

func (c *CommentService) UpdateComment(ctx context.Context, comment domain.Comment) error {
	ctx, span := c.tracer.Start(ctx, "UpdateComment")
	defer span.End()
	return c.commentRepository.UpdateComment(ctx, comment)
}

func (c *CommentService) DeleteComment(ctx context.Context, comment domain.Comment) error {
	ctx, span := c.tracer.Start(ctx, "DeleteComment")
	defer span.End()
	return c.commentRepository.DeleteComment(ctx, comment)
}
func (c *CommentService) CommentList(ctx context.Context) ([]domain.Comment, error) {
	ctx, span := c.tracer.Start(ctx, "CommentList")
	defer span.End()
	return c.commentRepository.CommentList(ctx)
}
