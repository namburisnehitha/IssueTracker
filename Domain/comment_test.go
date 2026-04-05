package domain

import (
	"testing"
)

func TestNewComment(t *testing.T) {
	issueid := "01"
	UserId := "bug"
	content := "need fix"
	id := "001"

	comment := NewComment(issueid, UserId, content, id)

	if comment.IssueId != issueid {
		t.Errorf("got %v,want %v", comment.IssueId, issueid)
	}
	if comment.UserId != UserId {
		t.Errorf("got %v,want %v", comment.UserId, UserId)
	}
	if comment.Content != content {
		t.Errorf("got %v,want %v", comment.Content, content)
	}
	if comment.Id != id {
		t.Errorf("got %v,want %v", comment.Id, id)
	}
}
