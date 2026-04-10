package domain

import (
	"testing"
	"time"
)

func TestNewIssue(t *testing.T) {

	id2 := "ISS-001"
	title2 := ""
	description2 := "no Title"
	_, err := NewIssue(id2, title2, description2)

	if err != ErrInvalidIssueData {
		t.Errorf("got %v,want %v", err, ErrInvalidIssueData)
	}

	id := "ISS-001"
	title := "Login broken"
	description := "Users cannot log in"

	issue, err := NewIssue(id, title, description)

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
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestStart(t *testing.T) {

	user := &User{Id: "user-123"}

	// Situation 1: open issue with assignee — should succeed

	issue1, err := NewIssue("001", "title", "desc")
	issue1.AssigneeId = "user-123"
	err = issue1.Start(user)

	// check err is nil

	if err != nil {
		t.Errorf("got %v , want %v", err, nil)
	}
	// check status is StatusInProgress

	if issue1.Status != StatusInProgress {
		t.Errorf("got %v,want %v", issue1.Status, StatusInProgress)
	}

	if user.Id != issue1.AssigneeId {
		t.Errorf("got %v,want %v", user.Id, issue1.AssigneeId)
	}

	// Situation 2: open issue, no assignee — should fail

	issue2, err := NewIssue("002", "title", "desc")
	err = issue2.Start(user)

	// check err is ErrIssueHasNoAssignee

	if err != ErrIssueHasNoAssignee {
		t.Errorf("got %v , want %v", err, ErrIssueHasNoAssignee)
	}

	// check status is still StatusOpen

	if issue2.Status != StatusOpen {
		t.Errorf("got %v , want %v", issue2.Status, StatusOpen)
	}

	// Situation 3: already in progress — should faild

	issue3, err := NewIssue("003", "title", "desc")
	issue3.Status = StatusInProgress
	t.Logf("issue3 status before Start: %v", issue3.Status)
	err = issue3.Start(user)

	// check err is ErrInvalidStateTransition
	if err != ErrInvalidStateTransition {
		t.Errorf("got %v, want %v", err, ErrInvalidStateTransition)
	}

	// Situation 4 : Wrong user id

	issue4, err := NewIssue("004", "title", "desc")
	issue4.AssigneeId = "user-123"
	user_1 := &User{Id: "user-456"}
	err = issue4.Start(user_1)

	if err != ErrUnauthorizedAction {
		t.Errorf("got %v,want %v", err, ErrUnauthorizedAction)
	}
}

func TestClose(t *testing.T) {

	user := &User{Id: "user-123"}

	//status is in progress
	issue, err := NewIssue("01", "title", "description")
	issue.AssigneeId = "user-123"
	issue.Status = StatusInProgress
	err = issue.Close(user)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if issue.Status != StatusClosed {
		t.Errorf("got %v,want %v", issue.Status, StatusClosed)
	}

	// status not in progress
	issue2, err := NewIssue("01", "title", "description")
	issue2.Status = StatusOpen
	issue2.AssigneeId = "user-123"
	err = issue2.Close(user)

	if err != ErrInvalidStateTransition {
		t.Errorf("got %v,want %v", err, ErrInvalidStateTransition)
	}

	issue3, err := NewIssue("01", "title", "description")
	issue3.Status = StatusClosed
	err = issue3.Close(user)

	if err != ErrInvalidStateTransition {
		t.Errorf("got %v,want %v", err, ErrInvalidStateTransition)
	}

	//Situation: when user is not assignee
	issue4, err := NewIssue("004", "title", "desc")
	issue4.AssigneeId = "user-123"
	issue4.Status = StatusInProgress
	user_1 := &User{Id: "user-456"}
	err = issue4.Close(user_1)

	if err != ErrUnauthorizedAction {
		t.Errorf("got %v,want %v", err, ErrUnauthorizedAction)
	}

}

func TestReOpen(t *testing.T) {
	user := &User{Id: "user-123"}
	//status is in closed
	issue, err := NewIssue("01", "title", "description")
	issue.AssigneeId = "user-123"
	issue.Status = StatusClosed
	err = issue.ReOpen(user)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if issue.Status != StatusOpen {
		t.Errorf("got %v,want %v", issue.Status, StatusOpen)
	}

	// status is not closed
	issue2, err := NewIssue("01", "title", "description")
	issue2.Status = StatusOpen
	issue2.AssigneeId = "user-123"
	err = issue2.ReOpen(user)

	if err != ErrInvalidStateTransition {
		t.Errorf("got %v,want %v", err, ErrInvalidStateTransition)
	}

	issue3, err := NewIssue("01", "title", "description")
	issue3.Status = StatusInProgress
	err = issue3.ReOpen(user)

	if err != ErrInvalidStateTransition {
		t.Errorf("got %v,want %v", err, ErrInvalidStateTransition)
	}
	// when user is not assignee
	issue4, err := NewIssue("004", "title", "desc")
	issue4.AssigneeId = "user-123"
	user_1 := &User{Id: "user-456"}
	issue4.Status = StatusClosed
	err = issue4.ReOpen(user_1)

	if err != ErrUnauthorizedAction {
		t.Errorf("got %v,want %v", err, ErrUnauthorizedAction)
	}

}

func TestAssignId(t *testing.T) {

	//not assigned
	issue, err := NewIssue("01", "Title", "desc")
	issue.AssigneeId = ""
	new_id := "06"
	err = issue.AssignIssue(new_id)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if issue.AssigneeId != new_id {
		t.Errorf("got %v,want %v", issue.AssigneeId, new_id)
	}
	//already assigned
	issue2, err := NewIssue("01", "Title", "desc")
	issue2.AssigneeId = "07"
	new_id = "06"
	err = issue2.AssignIssue(new_id)

	if err != ErrIssueAlreadyAssigned {
		t.Errorf("got %v,want %v", err, ErrIssueAlreadyAssigned)
	}
}
