package main

import (
	"fmt"
	"github.com/namburisnehitha/IssueTracker/handlers"
	"github.com/namburisnehitha/IssueTracker/internal/postgres"
	"github.com/namburisnehitha/IssueTracker/service"
	"log"
	"net/http"
)

func main() {

	db, err := postgres.NewDB("postgres://postgres:DAREALSQL@localhost:5432/issuetracker?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	issuerepo := postgres.NewPostgresIssueRepository(db)
	issueservice := service.NewIssueService(issuerepo)
	issuehandler := handlers.NewIssueHandler(issueservice)

	userrepo := postgres.NewPostgresUserRepository(db)
	userservice := service.NewUserService(userrepo)
	userhandler := handlers.NewUserHandler(userservice)

	labelrepo := postgres.NewPostgresLabelRepository(db)
	labelservice := service.NewLabelService(labelrepo)
	labelhandler := handlers.NewLabelHandler(labelservice)

	commentrepo := postgres.NewPostgresCommentRepository(db)
	commentservice := service.NewCommentService(commentrepo)
	commenthandler := handlers.NewCommentHandler(commentservice)

	activityrepo := postgres.NewPostgresActivityRepository(db)
	activityservice := service.NewActivityService(activityrepo)
	activityhandler := handlers.NewActivityHandler(activityservice)

	r := SetUpRoutes(issuehandler, userhandler, labelhandler, commenthandler, activityhandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
