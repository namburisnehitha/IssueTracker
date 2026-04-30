package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/domain"
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

func NewIssueHandler(issueService *service.IssueService) *IssueHandler {
	return &IssueHandler{
		issueService: issueService,
	}
}

func (i *IssueHandler) CreateIssue(w http.ResponseWriter, r *http.Request) {

	var ir CreateIssueRequest
	err := json.NewDecoder(r.Body).Decode(&ir)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ir.Id, err = i.issueService.CreateIssue(ir.Title, ir.Description)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, ir.Id)

}

func (i *IssueHandler) GetById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	issue, err := i.issueService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, issue)
}

func (i *IssueHandler) GetByTitle(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Query().Get("title")
	issue, err := i.issueService.GetByTitle(title)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, issue)
}

func (i *IssueHandler) GetByStatus(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")
	issue, err := i.issueService.GetByStatus(domain.IssueStatus(status))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, issue)
}

func (i *IssueHandler) UpdateIssue(w http.ResponseWriter, r *http.Request) {

	var ir CreateIssueRequest

	id := chi.URLParam(r, "id")
	issue, err := i.issueService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&ir)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	issue.Title = ir.Title
	issue.Description = ir.Description
	err = i.issueService.UpdateIssue(issue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, issue)

}

func (i *IssueHandler) DeleteIssue(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	issue, err := i.issueService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = i.issueService.DeleteIssue(issue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (i *IssueHandler) ListIssues(w http.ResponseWriter, r *http.Request) {

	issues, err := i.issueService.ListIssues()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, issues)
}
