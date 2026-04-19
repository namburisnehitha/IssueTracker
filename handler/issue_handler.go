package handlers

import (
	"encoding/json"
	"github.com/namburisnehitha/IssueTracker/service"
	"net/http"
)

type CreateIssueRequest struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type IssueHandler struct {
	issueService *service.IssueService
}

func NewIssueHandler(issueservice *service.IssueService) *IssueHandler {
	return &IssueHandler{
		issueService: issueservice,
	}
}

func (i *IssueHandler) CreateIssue(w http.ResponseWriter, r *http.Request) {

	var ir CreateIssueRequest
	err := json.NewDecoder(r.Body).Decode(&ir)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = i.issueService.CreateIssue(ir.Id, ir.Title, ir.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
