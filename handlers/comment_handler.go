package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/service"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type CreateCommentRequest struct {
	Id      string `json:"id"`
	IssueId string `json:"issueid"`
	UserId  string `json:"userid"`
	Content string `json:"content"`
}

type CommentHandler struct {
	commentService *service.CommentService
	tracer         trace.Tracer
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
		tracer:         otel.Tracer("comment-handler"),
	}
}

func (c *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {

	ctx, span := c.tracer.Start(r.Context(), "CreateComment")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var cr CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&cr)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	cr.Id, err = c.commentService.CreateComment(ctx, cr.IssueId, cr.UserId, cr.Content)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusCreated, cr.Id)

}

func (c *CommentHandler) GetById(w http.ResponseWriter, r *http.Request) {

	ctx, span := c.tracer.Start(r.Context(), "GetById")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	comment, err := c.commentService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusOK, comment)
}

func (c *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request) {

	ctx, span := c.tracer.Start(r.Context(), "GetComment")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	userid := r.URL.Query().Get("userid")
	issueid := r.URL.Query().Get("issueid")

	if userid != "" {
		comment, err := c.commentService.GetByUserId(ctx, userid)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, comment)
	} else if issueid != "" {

		comment, err := c.commentService.GetByIssueId(ctx, issueid)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, comment)
	} else {
		comments, err := c.commentService.CommentList(ctx)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, comments)
	}
}

func (c *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	ctx, span := c.tracer.Start(r.Context(), "UpdateComment")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var cr CreateCommentRequest

	id := chi.URLParam(r, "id")
	comment, err := c.commentService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}
	err = json.NewDecoder(r.Body).Decode(&cr)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	comment.Content = cr.Content
	err = c.commentService.UpdateComment(ctx, comment)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusOK, comment)

}

func (c *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	ctx, span := c.tracer.Start(r.Context(), "DeleteComment")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	comment, err := c.commentService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	err = c.commentService.DeleteComment(ctx, comment)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
