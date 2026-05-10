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
	tracer          trace.Tracer
}

func NewIssueService(issueRepository domain.IssueRepository) *IssueService {
	return &IssueService{
		issueRepository: issueRepository,
		tracer:          otel.Tracer("issue-service"),
	}
}

func (i *IssueService) CreateIssue(ctx context.Context, title string, description string, assigneeid string) (string, error) {

	ctx, span := i.tracer.Start(ctx, "CreateIssue")
	defer span.End()

	id := uuid.New().String()
	issue, err := domain.NewIssue(id, title, description, assigneeid)
	if err != nil {
		span.RecordError(err)
		return "", err
	}
	return id, i.issueRepository.Save(ctx, issue)
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
	ctx, span := i.tracer.Start(ctx, "DeleteIssue")
	defer span.End()
	return i.issueRepository.DeleteIssue(ctx, issue)
}

func (i *IssueService) ListIssues(ctx context.Context) ([]domain.Issue, error) {
	ctx, span := i.tracer.Start(ctx, "ListIssues")
	defer span.End()
	return i.issueRepository.ListIssues(ctx)
}
