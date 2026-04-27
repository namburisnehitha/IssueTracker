package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/domain"
	"github.com/namburisnehitha/IssueTracker/service"
)

type CreateActivityRequest struct {
	Id          string              `json:"id"`
	UserId      string              `json:"UserId"`
	IssueId     string              `json:"IssueId"`
	Description string              `json:"description"`
	Action      domain.ActivityType `json:"action"`
}

type ActivityHandler struct {
	activityService *service.ActivityService
}

func NewActivityHandler(activityService *service.ActivityService) *ActivityHandler {
	return &ActivityHandler{
		activityService: activityService,
	}
}

func (a *ActivityHandler) CreateNewActivity(w http.ResponseWriter, r *http.Request) {

	var ar CreateActivityRequest
	err := json.NewDecoder(r.Body).Decode(&ar)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.activityService.CreateActivity(ar.Id, ar.IssueId, ar.UserId, ar.Description, ar.Action)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (a *ActivityHandler) GetById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	activity, err := a.activityService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, activity)
}

func (a *ActivityHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {

	userid := r.URL.Query().Get("userid")
	activity, err := a.activityService.GetByUserId(userid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, activity)
}

func (a *ActivityHandler) GetByIssueId(w http.ResponseWriter, r *http.Request) {

	issueid := r.URL.Query().Get("issueid")
	activity, err := a.activityService.GetByIssueId(issueid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, activity)
}

func (a *ActivityHandler) GetByAction(w http.ResponseWriter, r *http.Request) {

	action := r.URL.Query().Get("action")
	activity, err := a.activityService.GetByAction(domain.ActivityType(action))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, activity)
}

func (a *ActivityHandler) ActivityList(w http.ResponseWriter, r *http.Request) {

	activities, err := a.activityService.ActivityList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, activities)
}
