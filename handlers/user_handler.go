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

type CreateUserRequest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserHandler struct {
	userservice *service.UserService
	tracer      trace.Tracer
}

func NewUserHandler(userservice *service.UserService) *UserHandler {
	return &UserHandler{
		userservice: userservice,
		tracer:      otel.Tracer("user-handler"),
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	ctx, span := u.tracer.Start(r.Context(), "CreateUser")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var ur CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&ur)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ur.Id, err = u.userservice.CreateUser(ctx, ur.Name, ur.UserName, ur.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, ur.Id)
}

func (u *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {

	ctx, span := u.tracer.Start(r.Context(), "GetById")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	user, err := u.userservice.GetById(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	ctx, span := u.tracer.Start(r.Context(), "GetUser")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	name := r.URL.Query().Get("name")
	role := r.URL.Query().Get("role")

	if name != "" {
		user, err := u.userservice.GetByName(ctx, name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJSON(w, http.StatusOK, user)
	} else if role != "" {
		user, err := u.userservice.GetByRole(ctx, domain.Roles(role))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJSON(w, http.StatusOK, user)
	} else {

		users, err := u.userservice.UserList(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJSON(w, http.StatusOK, users)
	}
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	ctx, span := u.tracer.Start(r.Context(), "UpdateUser")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var ur CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&ur)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := chi.URLParam(r, "id")
	user, err := u.userservice.GetById(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Name = ur.Name

	err = u.userservice.UpdateUser(ctx, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, user)

}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	ctx, span := u.tracer.Start(r.Context(), "DeleteUser")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	id := chi.URLParam(r, "id")
	user, err := u.userservice.GetById(ctx, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.userservice.DeleteUser(ctx, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
