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

func (i *IssueService) GetById(id string) (domain.Issue, error) {
	return i.issueRepository.GetById(id)
}

func (i *IssueService) GetByStatus(status domain.IssueStatus) ([]domain.Issue, error) {
	return i.issueRepository.GetByStatus(status)
}

func (i *IssueService) GetByTitle(title string) ([]domain.Issue, error) {
	return i.issueRepository.GetByTitle(title)
}

func (i *IssueService) UpdateIssue(issue domain.Issue) error {
	return i.issueRepository.UpdateIssue(issue)
}

func (i *IssueService) DeleteIssue(issue domain.Issue) error {
	return i.issueRepository.DeleteIssue(issue)
}

func (i *IssueService) ListIssues() ([]domain.Issue, error) {
	return i.issueRepository.ListIssues()
}
