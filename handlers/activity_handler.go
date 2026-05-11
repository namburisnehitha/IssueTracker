package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/domain"
	"github.com/namburisnehitha/IssueTracker/service"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
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
	tracer          trace.Tracer
}

func NewActivityHandler(activityService *service.ActivityService) *ActivityHandler {
	return &ActivityHandler{
		activityService: activityService,
		tracer:          otel.Tracer("activity-handler"),
	}
}

func (a *ActivityHandler) CreateNewActivity(w http.ResponseWriter, r *http.Request) {

	ctx, span := a.tracer.Start(r.Context(), "Createactivity")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var ar CreateActivityRequest
	err := json.NewDecoder(r.Body).Decode(&ar)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	ar.Id, err = a.activityService.CreateActivity(ctx, ar.IssueId, ar.UserId, ar.Description, ar.Action)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (a *ActivityHandler) GetById(w http.ResponseWriter, r *http.Request) {

	ctx, span := a.tracer.Start(r.Context(), "GetById")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	activity, err := a.activityService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusOK, activity)
}

func (a *ActivityHandler) Getactivity(w http.ResponseWriter, r *http.Request) {

	ctx, span := a.tracer.Start(r.Context(), "Getactivity")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	userid := r.URL.Query().Get("userid")
	issueid := r.URL.Query().Get("issueid")
	action := r.URL.Query().Get("action")

	if userid != "" {

		activity, err := a.activityService.GetByUserId(ctx, userid)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, activity)
	} else if issueid != "" {

		activity, err := a.activityService.GetByIssueId(ctx, issueid)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, activity)
	} else if action != "" {

		activity, err := a.activityService.GetByAction(ctx, domain.ActivityType(action))

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, activity)
	} else {

		activities, err := a.activityService.ActivityList(ctx)
		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, activities)
	}
}
