package domain

import "time"

type IssueStatus string

const (
	OPEN        IssueStatus = "OPEN"
	CLOSE       IssueStatus = "CLOSE"
	IN_PROGRESS IssueStatus = "IN_PROGRESS"
)

type issue struct {
	Id          int
	Title       string
	Description string
	Status      IssueStatus
	CreatedAt   time.Time
}
