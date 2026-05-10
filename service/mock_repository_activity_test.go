package service

import (
	"context"
	"testing"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type MockActivityRepository struct {
	activities map[string]domain.Activity
}

func (m *MockActivityRepository) Save(ctx context.Context, activity domain.Activity) error {
	m.activities[activity.Id] = activity
	return nil
}

func (m *MockActivityRepository) GetById(ctx context.Context, id string) (domain.Activity, error) {
	activity := m.activities[id]
	return activity, nil
}

func (m *MockActivityRepository) GetByUserId(ctx context.Context, userid string) ([]domain.Activity, error) {
	var result []domain.Activity
	for _, activity := range m.activities {
		if activity.UserId == userid {
			result = append(result, activity)
		}
	}
	return result, nil
}

func (m *MockActivityRepository) GetByIssueId(ctx context.Context, issueid string) ([]domain.Activity, error) {
	var result []domain.Activity
	for _, activity := range m.activities {
		if activity.IssueId == issueid {
			result = append(result, activity)
		}
	}
	return result, nil
}

func (m *MockActivityRepository) GetByAction(ctx context.Context, action domain.ActivityType) ([]domain.Activity, error) {
	var result []domain.Activity
	for _, activity := range m.activities {
		if activity.Action == action {
			result = append(result, activity)
		}
	}
	return result, nil
}

func (m *MockActivityRepository) ActivityList(ctx context.Context) ([]domain.Activity, error) {
	var result []domain.Activity
	for _, activity := range m.activities {
		result = append(result, activity)
	}
	return result, nil
}

func TestCreateActivity(t *testing.T) {
	id := "01"
	issueid := "02"
	userid := "03"
	description := "description"
	action := domain.UserAssigned
	repo := &MockActivityRepository{activities: map[string]domain.Activity{}}
	service := NewActivityService(repo)
	id, err := service.CreateActivity(context.Background(), issueid, userid, description, action)
	saved := repo.activities[id]

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if saved.Id != id {
		t.Errorf("got %v,want %v", saved.Id, id)
	}

	if saved.IssueId != issueid {
		t.Errorf("got %v,want %v", saved.IssueId, issueid)
	}

	if saved.UserId != userid {
		t.Errorf("got %v,want %v", saved.UserId, userid)
	}

	if saved.Description != description {
		t.Errorf("got %v,want %v", saved.Description, description)
	}

	if saved.Action != action {
		t.Errorf("got %v,want %v", saved.Action, action)
	}
}

func TestActivityGetById(t *testing.T) {
	id := "01"
	repo := &MockActivityRepository{activities: map[string]domain.Activity{}}
	service := NewActivityService(repo)
	repo.activities[id] = domain.Activity{Id: id}
	activity, err := service.GetById(context.Background(), id)

	if activity.Id != id {
		t.Errorf("got %v,want %v", activity.Id, id)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestActivityGetByIssueId(t *testing.T) {
	issueid := "001"
	issueid1 := "001"
	issueid2 := "002"
	issueid3 := "001"
	repo := &MockActivityRepository{activities: map[string]domain.Activity{}}
	service := NewActivityService(repo)
	repo.activities["01"] = domain.Activity{Id: "01", IssueId: issueid1}
	repo.activities["02"] = domain.Activity{Id: "02", IssueId: issueid2}
	repo.activities["03"] = domain.Activity{Id: "03", IssueId: issueid3}

	activities, err := service.GetByIssueId(context.Background(), issueid)

	for _, activity := range activities {
		if activity.IssueId != issueid {
			t.Errorf("got %v, want %v", activity.IssueId, issueid)
		}
	}

	if len(activities) != 2 {
		t.Errorf("got %v,want %v", len(activities), 2)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestActivityGetByUserId(t *testing.T) {
	userid := "001"
	userid1 := "001"
	userid2 := "002"
	userid3 := "001"
	repo := &MockActivityRepository{activities: map[string]domain.Activity{}}
	service := NewActivityService(repo)
	repo.activities["01"] = domain.Activity{Id: "01", UserId: userid1}
	repo.activities["02"] = domain.Activity{Id: "02", UserId: userid2}
	repo.activities["03"] = domain.Activity{Id: "03", UserId: userid3}

	activities, err := service.GetByUserId(context.Background(), userid)

	for _, activity := range activities {
		if activity.UserId != userid {
			t.Errorf("got %v, want %v", activity.UserId, userid)
		}
	}

	if len(activities) != 2 {
		t.Errorf("got %v,want %v", len(activities), 2)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestActivityList(t *testing.T) {
	repo := &MockActivityRepository{activities: map[string]domain.Activity{}}
	service := NewActivityService(repo)
	repo.activities["01"] = domain.Activity{Id: "01"}
	repo.activities["02"] = domain.Activity{Id: "02"}
	repo.activities["03"] = domain.Activity{Id: "03"}

	activities, err := service.ActivityList(context.Background())

	if len(activities) != 3 {
		t.Errorf("got %v,want %v", len(activities), 3)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}
