package service

import (
	"github.com/namburisnehitha/IssueTracker/domain"
)

type CommentService struct {
	commentRepository domain.CommentRepository
}

func NewCommentService(commentRepository domain.CommentRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (c *CommentService) CreateComment(issueid string, userid string, content string, id string) error {
	comment, err := domain.NewComment(issueid, userid, content, id)
	if err != nil {
		return err
	}
	return c.commentRepository.Save(comment)
}

func (c *CommentService) GetByIssueId(issueid string) ([]domain.Comment, error) {
	return c.commentRepository.GetByIssueId(issueid)
}

func (c *CommentService) GetById(id string) (domain.Comment, error) {
	return c.commentRepository.GetById(id)
}

func (c *CommentService) GetByUserId(userid string) ([]domain.Comment, error) {
	return c.commentRepository.GetByUserId(userid)
}

func (c *CommentService) UpdateComment(comment domain.Comment) error {
	return c.commentRepository.UpdateComment(comment)
}

func (c *CommentService) DeleteComment(comment domain.Comment) error {
	return c.commentRepository.DeleteComment(comment)
}
func (c *CommentService) CommentList() ([]domain.Comment, error) {
	return c.commentRepository.CommentList()
}
