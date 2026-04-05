package domain

import (
	"testing"
	"time"
)

func TestNewIssue(t *testing.T) {
	id := "ISS-001"
	title := "Login broken"
	description := "Users cannot log in"

	issue := NewIssue(id, title, description)

	if issue.Id != id {
		t.Errorf("got %v,want %v", issue.Id, id)
	}

	if issue.Title != title {
		t.Errorf("got %v,want %v", issue.Title, title)
	}

	if issue.Description != description {
		t.Errorf("got %v,want %v", issue.Description, description)
	}

	if issue.Status != StatusOpen {
		t.Errorf("got %v,want %v", issue.Status, StatusOpen)
	}

	if issue.CreatedAt.IsZero() {
		t.Errorf("got %v,want %v", issue.CreatedAt, time.Now())
	}
}

func TestStart(t *testing.T) {

	// Situation 1: open issue with assignee — should succeed

	issue1 := NewIssue("001", "title", "desc")
	issue1.AssigneeId = "user-123"
	err := issue1.Start()

	// check err is nil

	if err != nil {
		t.Errorf("got %v , want %v", err, nil)
	}
	// check status is StatusInProgress

	if issue1.Status != StatusInProgress {
		t.Errorf("got %v,want %v", issue1.Status, StatusInProgress)
	}

	// Situation 2: open issue, no assignee — should fail

	issue2 := NewIssue("002", "title", "desc")
	err = issue2.Start()

	// check err is ErrIssueHasNoAssignee

	if err != ErrIssueHasNoAssignee {
		t.Errorf("got %v , want %v", err, ErrIssueHasNoAssignee)
	}

	// check status is still StatusOpen

	if issue2.Status != StatusOpen {
		t.Errorf("got %v , want %v", issue2.Status, StatusOpen)
	}

	// Situation 3: already in progress — should faild

	issue3 := NewIssue("003", "title", "desc")
	issue3.Status = StatusInProgress
	t.Logf("issue3 status before Start: %v", issue3.Status)
	err = issue3.Start()

	// check err is ErrInvalidStateTransition
	if err != ErrInvalidStateTransition {
		t.Errorf("got %v, want %v", err, ErrInvalidStateTransition)
	}

}

func TestClose(t *testing.T) {
	//status is in progress
	issue := NewIssue("01", "title", "description")
	issue.AssigneeId = "67"
	issue.Status = StatusInProgress
	err := issue.Close()

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if issue.Status != StatusClosed {
		t.Errorf("got %v,want %v", issue.Status, StatusClosed)
	}

	// status not in progress
	issue2 := NewIssue("01", "title", "description")
	issue2.Status = StatusOpen
	issue2.AssigneeId = "6767"
	err = issue2.Close()

	if err != ErrInvalidStateTransition {
		t.Errorf("got %v,want %v", err, ErrInvalidStateTransition)
	}

	issue3 := NewIssue("01", "title", "description")
	issue3.Status = StatusClosed
	err = issue3.Close()

	if err != ErrInvalidStateTransition {
		t.Errorf("got %v,want %v", err, ErrInvalidStateTransition)
	}
}
