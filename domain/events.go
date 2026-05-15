package domain

import "context"

type DomainEvent struct {
	Type        ActivityType
	IssueId     string
	UserId      string
	Description string
	CommentId   string
	LabelId     string
}

type EventPublisher interface {
	Publish(ctx context.Context, event DomainEvent) error
}
