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

type CreateLabelRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Colour      string `json:"colour"`
}

type LabelHandler struct {
	labelService *service.LabelService
	tracer       trace.Tracer
}

func NewLabelHandler(labelSevice *service.LabelService) *LabelHandler {
	return &LabelHandler{
		labelService: labelSevice,
		tracer:       otel.Tracer("label-handler"),
	}
}

func (l *LabelHandler) CreateLabel(w http.ResponseWriter, r *http.Request) {

	ctx, span := l.tracer.Start(r.Context(), "CreateLabel")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var lr CreateLabelRequest
	err := json.NewDecoder(r.Body).Decode(&lr)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	lr.Id, err = l.labelService.CreateLabel(ctx, lr.Name, lr.Description, lr.Colour)
	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (l *LabelHandler) GetById(w http.ResponseWriter, r *http.Request) {

	ctx, span := l.tracer.Start(r.Context(), "GetById")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	label, err := l.labelService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusOK, label)
}

func (l *LabelHandler) GetLabel(w http.ResponseWriter, r *http.Request) {

	ctx, span := l.tracer.Start(r.Context(), "GetLabel")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	name := r.URL.Query().Get("name")
	colour := r.URL.Query().Get("colour")

	if name != "" {
		label, err := l.labelService.GetByName(ctx, name)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, label)

	} else if colour != "" {

		labels, err := l.labelService.GetByColour(ctx, colour)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, labels)

	} else {

		labels, err := l.labelService.LabelList(ctx)

		if err != nil {
			span.RecordError(err)
			http.Error(w, err.Error(), domainErrorToStatus(err))
			return
		}

		writeJSON(w, http.StatusOK, labels)
	}
}

func (l *LabelHandler) UpdateLabel(w http.ResponseWriter, r *http.Request) {

	ctx, span := l.tracer.Start(r.Context(), "UpdateLabel")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var lr CreateLabelRequest
	err := json.NewDecoder(r.Body).Decode(&lr)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	id := chi.URLParam(r, "id")
	label, err := l.labelService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	label.Colour = lr.Colour
	label.Name = lr.Name
	label.Description = lr.Description
	err = l.labelService.UpdateLabel(ctx, label)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	writeJSON(w, http.StatusOK, label)
}

func (l *LabelHandler) DeleteLabel(w http.ResponseWriter, r *http.Request) {

	ctx, span := l.tracer.Start(r.Context(), "DeleteLabel")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	label, err := l.labelService.GetById(ctx, id)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	err = l.labelService.DeleteLabel(ctx, label)

	if err != nil {
		span.RecordError(err)
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
