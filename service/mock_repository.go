package service

import (
	"github.com/namburisnehitha/IssueTracker/domain"
)

type MockIssueRepository struct {
	issues map[string]domain.Issue
}

func (m *MockIssueRepository) Save(issue domain.Issue) error {
	m.issues[issue.Id] = issue
	return nil
}

func (m *MockIssueRepository) GetById(id string) (domain.Issue, error) {
	issue := m.issues[id]
	return issue, nil
}

func (m *MockIssueRepository) GetByStatus(status domain.IssueStatus) ([]domain.Issue, error) {
	var result []domain.Issue
	for _, issue := range m.issues {
		if issue.Status == status {
			result = append(result, issue)
		}
	}
	return result, nil
}

func (m *MockIssueRepository) GetByTitle(title string) ([]domain.Issue, error) {
	var result []domain.Issue
	for _, issue := range m.issues {
		if issue.Title == title {
			result = append(result, issue)
		}
	}
	return result, nil
}

func (m *MockIssueRepository) UpdateIssue(issue domain.Issue) error {
	m.issues[issue.Id] = issue
	return nil
}

func (m *MockIssueRepository) DeleteIssue(issue domain.Issue) error {
	delete(m.issues, issue.Id)
	return nil
}

func (m *MockIssueRepository) ListIssues() ([]domain.Issue, error) {
	var result []domain.Issue
	for _, issue := range m.issues {
		result = append(result, issue)
	}
	return result, nil
}
