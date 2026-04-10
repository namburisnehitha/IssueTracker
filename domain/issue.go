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

func (i *Issue) Start(u *User) error {

	if i.Status != StatusOpen {
		return ErrInvalidStateTransition
	}

	if i.AssigneeId == "" {
		return ErrIssueHasNoAssignee
	}
	if u.Id != i.AssigneeId {
		return ErrUnauthorizedAction
	}

	i.Status = StatusInProgress
	return nil

}

func (i *Issue) Close(u *User) error {

	if i.Status != StatusInProgress {
		return ErrInvalidStateTransition
	}

	if u.Id != i.AssigneeId {
		return ErrUnauthorizedAction
	}
	i.Status = StatusClosed
	return nil
}

func (i *Issue) ReOpen(u *User) error {

	if i.Status != StatusClosed {
		return ErrInvalidStateTransition
	}
	if u.Id != i.AssigneeId {
		return ErrUnauthorizedAction
	}
	i.Status = StatusOpen
	return nil

}

func NewIssue(Id string, Title string, Description string) (Issue, error) {
	if Title == "" {
		return Issue{}, ErrInvalidIssueData
	}
	return Issue{
		Id:          Id,
		Title:       Title,
		Description: Description,
		Status:      StatusOpen,
		CreatedAt:   time.Now(),
	}, nil
}
