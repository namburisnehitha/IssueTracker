package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/handlers"
	"github.com/namburisnehitha/IssueTracker/internal/auth"
	"github.com/namburisnehitha/IssueTracker/internal/telemetry"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetUpRoutes(
	issueHandler *handlers.IssueHandler,
	userHandler *handlers.UserHandler,
	labelHandler *handlers.LabelHandler,
	commentHandler *handlers.CommentHandler,
	activityHandler *handlers.ActivityHandler,
	authHandler *handlers.AuthHandler,
	m *telemetry.Metrics,
	logger *slog.Logger,

) chi.Router {
	r := chi.NewRouter()
	r.Use(telemetry.MetricsMiddleware(logger, m))
	r.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware)
		r.Get("/issues", issueHandler.GetIssue)
		r.Post("/issues", issueHandler.CreateIssue)
		r.Get("/issues/{id}", issueHandler.GetById)
		r.Put("/issues/{id}", issueHandler.UpdateIssue)
		r.Delete("/issues/{id}", issueHandler.DeleteIssue)

		r.Get("/users", userHandler.GetUser)
		r.Get("/users/{id}", userHandler.GetById)
		r.Put("/users/{id}", userHandler.UpdateUser)
		r.Delete("/users/{id}", userHandler.DeleteUser)

		r.Get("/labels", labelHandler.GetLabel)
		r.Post("/labels", labelHandler.CreateLabel)
		r.Get("/labels/{id}", labelHandler.GetById)
		r.Put("/labels/{id}", labelHandler.UpdateLabel)
		r.Delete("/labels/{id}", labelHandler.DeleteLabel)
		r.Post("/issues/{id}/labels", labelHandler.AddLabelToIssue)
		r.Delete("/issues/{id}/labels/{labelId}", labelHandler.RemoveLabelFromIssue)

		r.Get("/comments", commentHandler.GetComment)
		r.Post("/comments", commentHandler.CreateComment)
		r.Get("/comments/{id}", commentHandler.GetById)
		r.Put("/comments/{id}", commentHandler.UpdateComment)
		r.Delete("/comments/{id}", commentHandler.DeleteComment)

		r.Get("/activities", activityHandler.GetActivity)
		r.Get("/activities/{id}", activityHandler.GetById)
	})

	r.Post("/login", authHandler.Login)
	r.Post("/users", userHandler.CreateUser)
	r.Handle("/metrics", promhttp.Handler())

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})

	return r

}
