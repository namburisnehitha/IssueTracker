package domain

import "time"

type IssueStatus string

const (
	StatusOpen       IssueStatus = "OPEN"
	StatusClosed     IssueStatus = "CLOSED"
	StatusInProgress IssueStatus = "IN_PROGRESS"
)

type Issue struct {
	Id          string
	Title       string
	Description string
	Status      IssueStatus
	CreatedAt   time.Time
	AssigneeId  string
	Labels      []Label
	Comments    []Comment
	Activities  []Activity
}

func (i *Issue) AssignIssue(Id string) error {
	if i.AssigneeId == "" {
		i.AssigneeId = Id
		return nil
	}
	return ErrIssueAlreadyAssigned
}

func (i *Issue) Start() error {

	if i.Status != StatusOpen {
		return ErrInvalidStateTransition
	}

	if i.AssigneeId == "" {
		return ErrIssueHasNoAssignee
	}

	i.Status = StatusInProgress
	return nil

}

func (i *Issue) Close() error {
	if i.Status == StatusInProgress {
		i.Status = StatusClosed
		return nil
	}
	return ErrInvalidStateTransition
}

func (i *Issue) ReOpen() error {
	if i.Status == StatusClosed {
		i.Status = StatusOpen
		return nil
	}
	return ErrInvalidStateTransition
}

func NewIssue(Id string, Title string, Description string) Issue {
	return Issue{
		Id:          Id,
		Title:       Title,
		Description: Description,
		Status:      StatusOpen,
		CreatedAt:   time.Now(),
	}
}
