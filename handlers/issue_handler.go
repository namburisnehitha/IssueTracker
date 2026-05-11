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

type CreateIssueRequest struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AssigneeId  string `json:"assignee_id"`
}

type IssueHandler struct {
	issueService *service.IssueService
	tracer       trace.Tracer
}

func NewIssueHandler(issueService *service.IssueService) *IssueHandler {
	return &IssueHandler{
		issueService: issueService,
		tracer:       otel.Tracer("issue-handler"),
	}
}

func (i *IssueHandler) CreateIssue(w http.ResponseWriter, r *http.Request) {

	ctx, span := i.tracer.Start(r.Context(), "CreateIssue")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var ir CreateIssueRequest
	err := json.NewDecoder(r.Body).Decode(&ir)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	ir.Id, err = i.issueService.CreateIssue(ctx, ir.Title, ir.Description, ir.AssigneeId)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusCreated, ir.Id)

}

func (i *IssueHandler) GetById(w http.ResponseWriter, r *http.Request) {

	ctx, span := i.tracer.Start(r.Context(), "GetById")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	issue, err := i.issueService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusOK, issue)
}

func (i *IssueHandler) GetIssue(w http.ResponseWriter, r *http.Request) {

	ctx, span := i.tracer.Start(r.Context(), "GetIssue")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	title := r.URL.Query().Get("title")
	status := r.URL.Query().Get("status")

	if title != "" {
		issue, err := i.issueService.GetByTitle(ctx, title)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, issue)
	} else if status != "" {

		issues, err := i.issueService.GetByStatus(ctx, domain.IssueStatus(status))

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, issues)
	} else {

		issues, err := i.issueService.ListIssues(ctx)
		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, issues)
	}
}

func (i *IssueHandler) UpdateIssue(w http.ResponseWriter, r *http.Request) {

	ctx, span := i.tracer.Start(r.Context(), "UpdateIssue")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var ir CreateIssueRequest

	id := chi.URLParam(r, "id")
	issue, err := i.issueService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}
	err = json.NewDecoder(r.Body).Decode(&ir)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	issue.Title = ir.Title
	issue.Description = ir.Description
	err = i.issueService.UpdateIssue(ctx, issue)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusOK, issue)

}

func (i *IssueHandler) DeleteIssue(w http.ResponseWriter, r *http.Request) {

	ctx, span := i.tracer.Start(r.Context(), "DeleteIssue")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	issue, err := i.issueService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	err = i.issueService.DeleteIssue(ctx, issue)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
