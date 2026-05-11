package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/internal/auth"
	"github.com/namburisnehitha/IssueTracker/service"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type AuthHandler struct {
	userService *service.UserService
	tracer      trace.Tracer
}

type CreateUserLoginInfoRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(userservice *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userservice,
		tracer:      otel.Tracer("auth-handler"),
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	ctx, span := ah.tracer.Start(r.Context(), "login")
	span.SetAttributes(semconv.HTTPRequestMethodKey.String(r.Method), semconv.HTTPRouteKey.String(chi.RouteContext(r.Context()).RoutePattern()))
	defer span.End()

	var ur *CreateUserLoginInfoRequest
	err := json.NewDecoder(r.Body).Decode(&ur)

	if err != nil {
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	user, err := ah.userService.GetByUserName(ctx, ur.UserName)

	if err != nil {
		http.Error(w, err.Error(), domainErrorToStatus(err))
		return
	}

	ok := auth.CheckPasswordHash(ur.Password, user.Password)
	if !ok {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.Id)

	if err != nil {
		http.Error(w, err.Error(), domainErrorToStatus(err))
	}

	writeJSON(w, http.StatusOK, token)

}
