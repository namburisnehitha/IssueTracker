package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/namburisnehitha/IssueTracker/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type LabelService struct {
	labelRepository domain.LabelRepository
	tracer          trace.Tracer
}

func NewLabelService(labelrepository domain.LabelRepository) *LabelService {
	return &LabelService{
		labelRepository: labelrepository,
		tracer:          otel.Tracer("label-service"),
	}
}

func (l *LabelService) CreateLabel(ctx context.Context, name string, description string, colour string) (string, error) {

	ctx, span := l.tracer.Start(ctx, "CreateLabel")
	defer span.End()

	id := uuid.New().String()
	label, err := domain.NewLabel(id, name, description, colour)
	if err != nil {
		return "", err
	}
	return id, l.labelRepository.Save(ctx, label)
}

func (l *LabelService) GetById(ctx context.Context, id string) (domain.Label, error) {
	ctx, span := l.tracer.Start(ctx, "GetById")
	defer span.End()
	return l.labelRepository.GetById(ctx, id)
}

func (l *LabelService) GetByName(ctx context.Context, name string) (domain.Label, error) {
	ctx, span := l.tracer.Start(ctx, "GetByName")
	defer span.End()
	return l.labelRepository.GetByName(ctx, name)
}

func (l *LabelService) GetByColour(ctx context.Context, colour string) ([]domain.Label, error) {
	ctx, span := l.tracer.Start(ctx, "GetByColour")
	defer span.End()
	return l.labelRepository.GetByColour(ctx, colour)
}

func (l *LabelService) UpdateLabel(ctx context.Context, label domain.Label) error {
	ctx, span := l.tracer.Start(ctx, "UpdateLabel")
	defer span.End()
	return l.labelRepository.UpdateLabel(ctx, label)
}

func (l *LabelService) DeleteLabel(ctx context.Context, label domain.Label) error {
	ctx, span := l.tracer.Start(ctx, "DeleteLabel")
	defer span.End()
	return l.labelRepository.DeleteLabel(ctx, label)
}

func (l *LabelService) LabelList(ctx context.Context) ([]domain.Label, error) {
	ctx, span := l.tracer.Start(ctx, "LabelList")
	defer span.End()
	return l.labelRepository.LabelList(ctx)
}
