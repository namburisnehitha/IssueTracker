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

func (i *IssueService) CreateIssue(id string, title string, description string) error {
	issue, err := domain.NewIssue(id, title, description)
	if err != nil {
		return err
	}
	return i.issueRepository.Save(issue)
}
