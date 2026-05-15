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
	publisher         domain.EventPublisher
	tracer            trace.Tracer
}

func NewCommentService(commentRepository domain.CommentRepository, publisher domain.EventPublisher) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
		publisher:         publisher,
		tracer:            otel.Tracer("comment-service"),
	}
}

func (c *CommentService) CreateComment(ctx context.Context, issueid string, content string) (string, error) {

	userID := domain.UserIDFromContext(ctx)

	ctx, span := c.tracer.Start(ctx, "CreateComment")
	defer span.End()

	id := uuid.New().String()
	comment, err := domain.NewComment(issueid, userID, content, id)
	if err != nil {
		return "", err
	}

	err = c.commentRepository.Save(ctx, comment)

	if err != nil {
		return "", err
	}

	c.publisher.Publish(ctx, domain.DomainEvent{
		Type:        domain.CommentAdded,
		IssueId:     issueid,
		UserId:      userID,
		Description: "create comment",
	})
	return id, nil

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
	userID := domain.UserIDFromContext(ctx)

	ctx, span := c.tracer.Start(ctx, "DeleteComment")
	defer span.End()
	err := c.commentRepository.DeleteComment(ctx, comment)
	if err != nil {
		return err
	}

	c.publisher.Publish(ctx, domain.DomainEvent{
		Type:        domain.CommentDeleted,
		UserId:      userID,
		IssueId:     comment.IssueId,
		Description: "delete comment",
	})
	return nil
}
func (c *CommentService) CommentList(ctx context.Context) ([]domain.Comment, error) {
	ctx, span := c.tracer.Start(ctx, "CommentList")
	defer span.End()
	return c.commentRepository.CommentList(ctx)
}
