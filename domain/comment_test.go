package domain

import (
	"testing"
)

func TestNewComment(t *testing.T) {
	issueid2 := "01"
	UserId2 := "01"
	content2 := ""
	id2 := "colour"
	_, err := NewComment(issueid2, UserId2, content2, id2)

	if err != ErrInvalidCommentData {
		t.Fatalf("got %v,want %v", ErrInvalidCommentData, err)
	}

	issueid := "01"
	UserId := "bug"
	content := "need fix"
	id := "001"

	comment, err := NewComment(issueid, UserId, content, id)

	if err != nil {
		t.Fatalf("failed to create comment: %v", err)
	}

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
