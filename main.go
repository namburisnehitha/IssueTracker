package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/namburisnehitha/IssueTracker/handlers"
	"github.com/namburisnehitha/IssueTracker/internal/postgres"
	"github.com/namburisnehitha/IssueTracker/internal/telemetry"
	"github.com/namburisnehitha/IssueTracker/service"
	"go.opentelemetry.io/otel"
)

func main() {

	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	db, err := postgres.NewDB(dbURL)
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

	authHandler := handlers.NewAuthHandler(userservice)

	CleanupTracer, err := telemetry.InitTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer CleanupTracer()

	CleanupMeter, err := telemetry.InitMeter()
	if err != nil {
		log.Fatal(err)
	}
	defer CleanupMeter()

	m, err := telemetry.NewMetrics(otel.Meter("issue-tracker"))
	if err != nil {
		log.Fatal(err)
	}

	logger := telemetry.InitLogger()

	r := SetUpRoutes(issuehandler, userhandler, labelhandler, commenthandler, activityhandler, authHandler, &m, logger)
	server := &http.Server{Addr: ":8080", Handler: r}

	go func() {
		fmt.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)

}
