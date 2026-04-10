package domain

import (
	"time"
)

type Comment struct {
	IssueId   string
	UserId    string
	Content   string
	CreatedAt time.Time
	Id        string
}

func NewComment(IssueId string, UserId string, Content string, Id string) Comment {
	return Comment{
		IssueId:   IssueId,
		UserId:    UserId,
		Content:   Content,
		CreatedAt: time.Now(),
		Id:        Id,
	}
}
