package domain

type IssueRepository interface {
	NewIssue(issue Issue) error
	GetById(id string) (Issue, error)
	GetByStatus(status []IssueStatus) ([]Issue, error)
	GetByTitle(title string) ([]Issue, error)
	UpdateIssue(issue Issue) error
	DeleteIssue(issue Issue) error
	ListIssues() ([]Issue, error)
}
