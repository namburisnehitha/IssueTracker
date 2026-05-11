package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/namburisnehitha/IssueTracker/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type ActivityService struct {
	activityRepository domain.ActivityRepository
	tracer             trace.Tracer
}

func NewActivityService(activityRepository domain.ActivityRepository) *ActivityService {
	return &ActivityService{
		activityRepository: activityRepository,
		tracer:             otel.Tracer("activity-service"),
	}
}

func (a *ActivityService) CreateActivity(ctx context.Context, issueid string, userid string, description string, action domain.ActivityType) (string, error) {

	ctx, span := a.tracer.Start(ctx, "CreateActivity")
	defer span.End()

	id := uuid.New().String()
	activity, err := domain.NewActivity(id, issueid, userid, description, action)
	if err != nil {
		return "", err
	}
	return id, a.activityRepository.Save(ctx, activity)

}

func (a *ActivityService) GetById(ctx context.Context, id string) (domain.Activity, error) {
	ctx, span := a.tracer.Start(ctx, "GetById")
	defer span.End()
	return a.activityRepository.GetById(ctx, id)
}

func (a *ActivityService) GetByUserId(ctx context.Context, userid string) ([]domain.Activity, error) {
	ctx, span := a.tracer.Start(ctx, "GetByUserId")
	defer span.End()
	return a.activityRepository.GetByUserId(ctx, userid)
}

func (a *ActivityService) GetByIssueId(ctx context.Context, issueid string) ([]domain.Activity, error) {
	ctx, span := a.tracer.Start(ctx, "GetByIssueId")
	defer span.End()
	return a.activityRepository.GetByIssueId(ctx, issueid)
}

func (a *ActivityService) GetByAction(ctx context.Context, action domain.ActivityType) ([]domain.Activity, error) {
	ctx, span := a.tracer.Start(ctx, "GetByAction")
	defer span.End()
	return a.activityRepository.GetByAction(ctx, action)
}

func (a *ActivityService) ActivityList(ctx context.Context) ([]domain.Activity, error) {
	ctx, span := a.tracer.Start(ctx, "ActivityList")
	defer span.End()
	return a.activityRepository.ActivityList(ctx)
}
