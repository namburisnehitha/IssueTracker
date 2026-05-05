package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/handlers"
)

func SetUpRoutes(
	issueHandler *handlers.IssueHandler,
	userHandler *handlers.UserHandler,
	labelHandler *handlers.LabelHandler,
	commentHandler *handlers.CommentHandler,
	activityHandler *handlers.ActivityHandler,
) chi.Router {
	r := chi.NewRouter()

	r.Get("/issues", issueHandler.GetIssue)
	r.Post("/issues", issueHandler.CreateIssue)
	r.Get("/issues/{id}", issueHandler.GetById)
	r.Put("/issues/{id}", issueHandler.UpdateIssue)
	r.Delete("/issues/{id}", issueHandler.DeleteIssue)

	r.Get("/users", userHandler.GetUser)
	r.Post("/users", userHandler.CreateUser)
	r.Get("/users/{id}", userHandler.GetById)
	r.Put("/users/{id}", userHandler.UpdateUser)
	r.Delete("/users/{id}", userHandler.DeleteUser)

	r.Get("/labels", labelHandler.GetLabel)
	r.Post("/labels", labelHandler.CreateLabel)
	r.Get("/labels/{id}", labelHandler.GetById)
	r.Put("/labels/{id}", labelHandler.UpdateLabel)
	r.Delete("/labels/{id}", labelHandler.DeleteLabel)

	r.Get("/comments", commentHandler.GetComment)
	r.Post("/comments", commentHandler.CreateComment)
	r.Get("/comments/{id}", commentHandler.GetById)
	r.Put("/comments/{id}", commentHandler.UpdateComment)
	r.Delete("/comments/{id}", commentHandler.DeleteComment)

	r.Get("/activities", activityHandler.Getactivity)
	r.Post("/activities", activityHandler.CreateNewActivity)
	return r
}
