package service

import (
	"testing"

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

func TestCreateIssue(t *testing.T) {
	id := "01"
	title := "Create Isuue"
	description := "create the issue"
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service := NewIssueService(repo)
	err := service.CreateIssue(id, title, description)
	saved := repo.issues[id]

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if saved.Id != id {
		t.Errorf("got %v,want %v", saved.Id, id)
	}

	if saved.Title != title {
		t.Errorf("got %v,want %v", saved.Title, title)
	}
	if saved.Description != description {
		t.Errorf("got %v,want %v", saved.Description, description)
	}
}

func TestGetById(t *testing.T) {
	id := "01"
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service := NewIssueService(repo)
	repo.issues[id] = domain.Issue{Id: id, Title: "test"}
	issue, err := service.GetById(id)

	if issue.Id != id {
		t.Errorf("got %v,want %v", issue.Id, id)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

}
