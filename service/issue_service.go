package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/namburisnehitha/IssueTracker/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type IssueService struct {
	issueRepository domain.IssueRepository
	publisher       domain.EventPublisher
	tracer          trace.Tracer
}

func NewIssueService(issueRepository domain.IssueRepository, publisher domain.EventPublisher) *IssueService {
	return &IssueService{
		issueRepository: issueRepository,
		publisher:       publisher,
		tracer:          otel.Tracer("issue-service"),
	}
}

func (i *IssueService) CreateIssue(ctx context.Context, title string, description string, assigneeid string) (string, error) {

	userID := domain.UserIDFromContext(ctx)

	ctx, span := i.tracer.Start(ctx, "CreateIssue")
	defer span.End()

	id := uuid.New().String()
	issue, err := domain.NewIssue(id, title, description, assigneeid)
	if err != nil {
		span.RecordError(err)
		return "", err
	}
	err = i.issueRepository.Save(ctx, issue)

	if err != nil {
		return "", err
	}

	i.publisher.Publish(ctx, domain.DomainEvent{
		Type:        domain.IssueCreated,
		IssueId:     id,
		UserId:      userID,
		Description: "issue created",
	})
	return id, nil

}

func (i *IssueService) GetById(ctx context.Context, id string) (domain.Issue, error) {
	ctx, span := i.tracer.Start(ctx, "GetById")
	defer span.End()
	return i.issueRepository.GetById(ctx, id)
}

func (i *IssueService) GetByStatus(ctx context.Context, status domain.IssueStatus) ([]domain.Issue, error) {
	ctx, span := i.tracer.Start(ctx, "GetByStatus")
	defer span.End()
	return i.issueRepository.GetByStatus(ctx, status)
}

func (i *IssueService) GetByTitle(ctx context.Context, title string) ([]domain.Issue, error) {
	ctx, span := i.tracer.Start(ctx, "GetByTitle")
	defer span.End()
	return i.issueRepository.GetByTitle(ctx, title)
}

func (i *IssueService) UpdateIssue(ctx context.Context, issue domain.Issue) error {
	ctx, span := i.tracer.Start(ctx, "UpdateIssue")
	defer span.End()
	return i.issueRepository.UpdateIssue(ctx, issue)
}

func (i *IssueService) DeleteIssue(ctx context.Context, issue domain.Issue) error {

	userID := domain.UserIDFromContext(ctx)

	ctx, span := i.tracer.Start(ctx, "DeleteIssue")
	defer span.End()
	err := i.issueRepository.DeleteIssue(ctx, issue)

	if err != nil {
		return err
	}

	i.publisher.Publish(ctx, domain.DomainEvent{
		Type:        domain.IssueDeleted,
		IssueId:     issue.Id,
		UserId:      userID,
		Description: "issue deleted",
	})
	return nil
}

func (i *IssueService) ListIssues(ctx context.Context) ([]domain.Issue, error) {
	ctx, span := i.tracer.Start(ctx, "ListIssues")
	defer span.End()
	return i.issueRepository.ListIssues(ctx)
}
