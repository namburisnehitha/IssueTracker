package domain

import (
	"time"
)

type Comment struct {
	IssueId   string
	UserId    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        string
}

func NewComment(IssueId string, UserId string, Content string, Id string) (Comment, error) {
	if Content == "" {
		return Comment{}, ErrInvalidCommentData
	}
	return Comment{
		IssueId:   IssueId,
		UserId:    UserId,
		Content:   Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Id:        Id,
	}, nil
}
