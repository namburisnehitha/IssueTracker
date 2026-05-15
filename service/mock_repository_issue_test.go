package service

import (
	"context"
	"testing"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type MockIssueRepository struct {
	issues map[string]domain.Issue
}

type MockEventPublisher struct{}

func (m *MockEventPublisher) Publish(ctx context.Context, event domain.DomainEvent) error {
	return nil
}

func (m *MockIssueRepository) Save(ctx context.Context, issue domain.Issue) error {
	m.issues[issue.Id] = issue
	return nil
}

func (m *MockIssueRepository) GetById(ctx context.Context, id string) (domain.Issue, error) {
	issue := m.issues[id]
	return issue, nil
}

func (m *MockIssueRepository) GetByStatus(ctx context.Context, status domain.IssueStatus) ([]domain.Issue, error) {
	var result []domain.Issue
	for _, issue := range m.issues {
		if issue.Status == status {
			result = append(result, issue)
		}
	}
	return result, nil
}

func (m *MockIssueRepository) GetByTitle(ctx context.Context, title string) ([]domain.Issue, error) {
	var result []domain.Issue
	for _, issue := range m.issues {
		if issue.Title == title {
			result = append(result, issue)
		}
	}
	return result, nil
}

func (m *MockIssueRepository) UpdateIssue(ctx context.Context, issue domain.Issue) error {
	m.issues[issue.Id] = issue
	return nil
}

func (m *MockIssueRepository) DeleteIssue(ctx context.Context, issue domain.Issue) error {
	delete(m.issues, issue.Id)
	return nil
}

func (m *MockIssueRepository) ListIssues(ctx context.Context) ([]domain.Issue, error) {
	var result []domain.Issue
	for _, issue := range m.issues {
		result = append(result, issue)
	}
	return result, nil
}

func TestCreateIssue(t *testing.T) {
	title := "Create Isuue"
	description := "create the issue"
	assignee_id := "001"
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}

	service := NewIssueService(repo, &MockEventPublisher{})

	id, err := service.CreateIssue(context.Background(), title, description, assignee_id)
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
	if saved.AssigneeId != assignee_id {
		t.Errorf("got %v,want %v", saved.AssigneeId, assignee_id)
	}
}

func TestIsssueGetById(t *testing.T) {
	id := "01"
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service := NewIssueService(repo, &MockEventPublisher{})
	repo.issues[id] = domain.Issue{Id: id, Title: "test"}
	issue, err := service.GetById(context.Background(), id)

	if issue.Id != id {
		t.Errorf("got %v,want %v", issue.Id, id)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

}

func TestIssueGetByStatus(t *testing.T) {
	//only one issue
	status := domain.StatusOpen
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service := NewIssueService(repo, &MockEventPublisher{})
	repo.issues["01"] = domain.Issue{Id: "01", Title: "test", Status: status}
	issue, err := service.GetByStatus(context.Background(), status)

	if issue[0].Status != status {
		t.Errorf("got %v,want %v", issue[0].Status, status)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	//mulitple issues(actual use case)
	status0 := domain.StatusOpen
	status1 := domain.StatusOpen
	status2 := domain.StatusClosed
	status3 := domain.StatusOpen
	repo1 := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service1 := NewIssueService(repo1, &MockEventPublisher{})
	repo1.issues["01"] = domain.Issue{Id: "01", Title: "test", Status: status1}
	repo1.issues["02"] = domain.Issue{Id: "02", Title: "test", Status: status2}
	repo1.issues["03"] = domain.Issue{Id: "03", Title: "test", Status: status3}
	issues, err := service1.GetByStatus(context.Background(), status0)
	for _, issue := range issues {
		if issue.Status != status0 {
			t.Errorf("got %v,want %v", issue.Status, status)
		}
	}
	if len(issues) != 2 {
		t.Errorf("got %v, want %v", len(issues), 2)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

}

func TestIssueGetByTitle(t *testing.T) {
	//with one issue
	title := "test"
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service := NewIssueService(repo, &MockEventPublisher{})
	repo.issues["01"] = domain.Issue{Id: "01", Title: title}
	issue, err := service.GetByTitle(context.Background(), title)

	if issue[0].Title != title {
		t.Errorf("got %v,want %v", issue[0].Title, title)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
	//with multiple issues(actual use case)
	title0 := "test0"
	title1 := "test1"
	title2 := "test0"
	title3 := "test0"

	repo1 := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service1 := NewIssueService(repo1, &MockEventPublisher{})
	repo1.issues["01"] = domain.Issue{Id: "01", Title: title1}
	repo1.issues["02"] = domain.Issue{Id: "02", Title: title2}
	repo1.issues["03"] = domain.Issue{Id: "03", Title: title3}
	issues, err := service1.GetByTitle(context.Background(), title0)
	for _, issue := range issues {
		if issue.Title != title0 {
			t.Errorf("got %v,want %v", issue.Title, title)
		}
	}
	if len(issues) != 2 {
		t.Errorf("got %v, want %v", len(issues), 2)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

}

func TestUpdateIssue(t *testing.T) {
	title := "new"
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service := NewIssueService(repo, &MockEventPublisher{})
	repo.issues["01"] = domain.Issue{Id: "01", Title: "old"}
	issue := domain.Issue{Id: "01", Title: title}
	err := service.UpdateIssue(context.Background(), issue)
	updated := repo.issues["01"]

	if updated.Title != title {
		t.Errorf("got %v,want %v", updated.Title, title)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestDeleteIssue(t *testing.T) {
	id := "01"
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service := NewIssueService(repo, &MockEventPublisher{})
	repo.issues[id] = domain.Issue{Id: id}
	issue := domain.Issue{Id: id}
	err := service.DeleteIssue(context.Background(), issue)

	_, exists := repo.issues["01"]

	if exists {
		t.Errorf("issue was not deleted")
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestListIssue(t *testing.T) {
	repo := &MockIssueRepository{issues: map[string]domain.Issue{}}
	service1 := NewIssueService(repo, &MockEventPublisher{})
	repo.issues["01"] = domain.Issue{Id: "01", Title: "test", Status: domain.StatusOpen}
	repo.issues["02"] = domain.Issue{Id: "02", Title: "test", Status: domain.StatusOpen}
	repo.issues["03"] = domain.Issue{Id: "03", Title: "test", Status: domain.StatusOpen}
	issues, err := service1.ListIssues(context.Background())
	if len(issues) != 3 {
		t.Errorf("got %v, want %v", len(issues), 3)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}
