package service

import "github.com/namburisnehitha/IssueTracker/domain"

type IssueService struct {
	issueRepository domain.IssueRepository
}

func NewIssueService(issueRepository domain.IssueRepository) IssueService {
	return IssueService{
		issueRepository: issueRepository,
	}
}

func (i *IssueService) CreateIssue(Id string, Title string, Description string) error {
	domain.NewIssue(Id, Title, Description)
	NewIssueService(i.issueRepository)

}
