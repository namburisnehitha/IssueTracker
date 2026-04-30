package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/service"
	"net/http"
)

type CreateCommentRequest struct {
	Id      string `json:"id"`
	IssueId string `json:"issueid"`
	UserId  string `json:"userid"`
	Content string `json:"content"`
}

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (c *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {

	var cr CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&cr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cr.Id, err = c.commentService.CreateComment(cr.IssueId, cr.UserId, cr.Content)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, cr.Id)

}

func (c *CommentHandler) GetById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	comment, err := c.commentService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, comment)
}

func (c *CommentHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {

	userid := r.URL.Query().Get("userid")
	comment, err := c.commentService.GetByUserId(userid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, comment)
}

func (c *CommentHandler) GetByIssueId(w http.ResponseWriter, r *http.Request) {

	issueid := r.URL.Query().Get("issueid")
	comment, err := c.commentService.GetByIssueId(issueid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, comment)
}

func (c *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	var cr CreateCommentRequest

	id := chi.URLParam(r, "id")
	comment, err := c.commentService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&cr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.Content = cr.Content
	err = c.commentService.UpdateComment(comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, comment)

}

func (c *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	comment, err := c.commentService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.commentService.DeleteComment(comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *CommentHandler) CommentList(w http.ResponseWriter, r *http.Request) {

	comments, err := c.commentService.CommentList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, comments)
}
