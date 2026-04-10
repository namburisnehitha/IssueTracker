package domain

import (
	"testing"
	"time"
)

func TestNewActivity(t *testing.T) {
	userid := "01"
	issueid := "001"
	description := "added comment"
	action := CommentAdded

	act := NewActivity(issueid, userid, description, action)

	if act.IssueId != issueid {
		t.Errorf("got %v,want %v", act.IssueId, issueid)
	}

	if act.UserId != userid {
		t.Errorf("got %v,want %v", act.UserId, userid)
	}

	if act.Description != description {
		t.Errorf("got %v,want %v", act.Description, description)
	}

	if act.Action != action {
		t.Errorf("got %v,want %v", act.Action, action)
	}

	if act.CreatedAt.IsZero() {
		t.Errorf("got %v,want %v", act.CreatedAt, time.Now())
	}
}
