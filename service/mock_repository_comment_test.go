package service

import (
	"context"
	"testing"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type MockCommentRepository struct {
	comments map[string]domain.Comment
}

func (m *MockCommentRepository) Save(ctx context.Context, comment domain.Comment) error {
	m.comments[comment.Id] = comment
	return nil
}

func (m *MockCommentRepository) GetById(ctx context.Context, id string) (domain.Comment, error) {
	comment := m.comments[id]
	return comment, nil
}

func (m *MockCommentRepository) GetByUserId(ctx context.Context, userid string) ([]domain.Comment, error) {
	var result []domain.Comment
	for _, comment := range m.comments {
		if comment.UserId == userid {
			result = append(result, comment)
		}
	}
	return result, nil
}

func (m *MockCommentRepository) GetByIssueId(ctx context.Context, issueid string) ([]domain.Comment, error) {
	var result []domain.Comment
	for _, comment := range m.comments {
		if comment.IssueId == issueid {
			result = append(result, comment)
		}
	}
	return result, nil
}

func (m *MockCommentRepository) UpdateComment(ctx context.Context, comment domain.Comment) error {
	m.comments[comment.Id] = comment
	return nil
}

func (m *MockCommentRepository) DeleteComment(ctx context.Context, comment domain.Comment) error {
	delete(m.comments, comment.Id)
	return nil
}

func (m *MockCommentRepository) CommentList(ctx context.Context) ([]domain.Comment, error) {
	var result []domain.Comment
	for _, comment := range m.comments {
		result = append(result, comment)
	}
	return result, nil
}

func TestCreateNewComment(t *testing.T) {
	issueid := "01"
	content := "Create new comment"
	repo := &MockCommentRepository{comments: map[string]domain.Comment{}}
	service := NewCommentService(repo, &MockEventPublisher{})
	id, err := service.CreateComment(context.Background(), issueid, content)
	saved := repo.comments[id]
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if saved.IssueId != issueid {
		t.Errorf("got %v,want %v", saved.IssueId, issueid)
	}

	if saved.Content != content {
		t.Errorf("got %v,want %v", saved.Content, content)
	}
	if saved.Id != id {
		t.Errorf("got %v,want %v", saved.Id, id)
	}
}

func TestCommentGetById(t *testing.T) {
	id := "01"
	repo := &MockCommentRepository{comments: map[string]domain.Comment{}}
	service := NewCommentService(repo, &MockEventPublisher{})
	repo.comments[id] = domain.Comment{Id: id}
	comment, err := service.GetById(context.Background(), id)
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
	if comment.Id != id {
		t.Errorf("got %v,want %v", comment.Id, id)
	}
}

func TestCommentGetByIssueId(t *testing.T) {
	issueid1 := "10"
	issueid2 := "10"
	issueid3 := "30"

	repo := &MockCommentRepository{comments: map[string]domain.Comment{}}
	service := NewCommentService(repo, &MockEventPublisher{})
	repo.comments["01"] = domain.Comment{Id: "01", IssueId: issueid1}
	repo.comments["02"] = domain.Comment{Id: "02", IssueId: issueid2}
	repo.comments["03"] = domain.Comment{Id: "03", IssueId: issueid3}
	comments, err := service.GetByIssueId(context.Background(), issueid1)

	for _, comment := range comments {
		if comment.IssueId != issueid1 {
			t.Errorf("got %v,want %v", comment.IssueId, issueid1)
		}
	}

	if len(comments) != 2 {
		t.Errorf("got %v,want %v", len(comments), 2)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

}

func TestCommentGetByUserId(t *testing.T) {
	userid1 := "10"
	userid2 := "10"
	userid3 := "30"

	repo := &MockCommentRepository{comments: map[string]domain.Comment{}}
	service := NewCommentService(repo, &MockEventPublisher{})
	repo.comments["01"] = domain.Comment{Id: "01", UserId: userid1}
	repo.comments["02"] = domain.Comment{Id: "02", UserId: userid2}
	repo.comments["03"] = domain.Comment{Id: "03", UserId: userid3}
	comments, err := service.GetByUserId(context.Background(), userid1)

	for _, comment := range comments {
		if comment.UserId != userid1 {
			t.Errorf("got %v,want %v", comment.UserId, userid1)
		}
	}

	if len(comments) != 2 {
		t.Errorf("got %v,want %v", len(comments), 2)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

}

func TestUpdateComment(t *testing.T) {
	content := "New Comment"
	id := "1"
	repo := &MockCommentRepository{comments: map[string]domain.Comment{}}
	service := NewCommentService(repo, &MockEventPublisher{})
	repo.comments[id] = domain.Comment{Id: id, Content: "old comment"}
	comment := domain.Comment{Id: id, Content: content}
	err := service.UpdateComment(context.Background(), comment)
	updated := repo.comments[id]

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if updated.Content != content {
		t.Errorf("got %v,want %v", updated.Content, content)
	}

}

func TestDeleteComment(t *testing.T) {
	id := "01"
	repo := &MockCommentRepository{comments: map[string]domain.Comment{}}
	service := NewCommentService(repo, &MockEventPublisher{})
	repo.comments[id] = domain.Comment{Id: id, Content: "comment"}
	comment := domain.Comment{Id: id, Content: "comment"}
	err := service.DeleteComment(context.Background(), comment)
	_, exists := repo.comments["01"]

	if exists {
		t.Errorf("comment was not deleted")
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

}

func TestListComments(t *testing.T) {

	repo := &MockCommentRepository{comments: map[string]domain.Comment{}}
	service := NewCommentService(repo, &MockEventPublisher{})
	repo.comments["01"] = domain.Comment{Id: "01", Content: "comment"}
	repo.comments["02"] = domain.Comment{Id: "02", Content: "comment"}
	repo.comments["03"] = domain.Comment{Id: "03", Content: "comment"}
	comments, err := service.CommentList(context.Background())

	if len(comments) != 3 {
		t.Errorf("got %v,want %v", len(comments), 3)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}
