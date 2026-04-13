package service

import (
	"github.com/namburisnehitha/IssueTracker/domain"
	"testing"
)

type MockCommentRepository struct {
	comments map[string]domain.Comment
}

func (m *MockCommentRepository) Save(comment domain.Comment) error {
	m.comments[comment.Id] = comment
	return nil
}

func (m *MockCommentRepository) GetById(id string) (domain.Comment, error) {
	comment := m.comments[id]
	return comment, nil
}

func (m *MockCommentRepository) GetByUserId(userid string) ([]domain.Comment, error) {
	var result []domain.Comment
	for _, comment := range m.comments {
		if comment.UserId == userid {
			result = append(result, comment)
		}
	}
	return result, nil
}

func (m *MockCommentRepository) GetByIssueId(issueid string) ([]domain.Comment, error) {
	var result []domain.Comment
	for _, comment := range m.comments {
		if comment.IssueId == issueid {
			result = append(result, comment)
		}
	}
	return result, nil
}

func (m *MockCommentRepository) UpdateComment(comment domain.Comment) error {
	m.comments[comment.Id] = comment
	return nil
}

func (m *MockCommentRepository) DeleteComment(comment domain.Comment) error {
	delete(m.comments, comment.Id)
	return nil
}

func (m *MockCommentRepository) CommentList() ([]domain.Comment, error) {
	var result []domain.Comment
	for _, comment := range m.comments {
		result = append(result, comment)
	}
	return result, nil
}

func TestCreateNewComment(t *testing.T) {
	issueid := "01"
	userid := "01"
	content := "Create a ne comment"
	id := "1"
	repo := &MockCommentRepository{comments: map[string]domain.Comment{}}
	service := NewCommentService(repo)
	err := service.CreateComment(issueid, userid, content, id)
	saved := repo.comments[id]
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if saved.IssueId != issueid {
		t.Errorf("got %v,want %v", saved.IssueId, issueid)
	}

	if saved.UserId != userid {
		t.Errorf("got %v,want %v", saved.UserId, userid)
	}
	if saved.Content != content {
		t.Errorf("got %v,want %v", saved.Content, content)
	}
	if saved.Id != id {
		t.Errorf("got %v,want %v", saved.Id, id)
	}
}
