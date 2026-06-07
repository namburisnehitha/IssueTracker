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

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	runMigrations(dbURL)

	db, err := postgres.NewDB(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	runMigrations(dbURL)

	activityrepo := postgres.NewPostgresActivityRepository(db)
	activityservice := service.NewActivityService(activityrepo)
	activityhandler := handlers.NewActivityHandler(activityservice)

	issuerepo := postgres.NewPostgresIssueRepository(db)
	issueservice := service.NewIssueService(issuerepo, activityservice)
	issuehandler := handlers.NewIssueHandler(issueservice)

	userrepo := postgres.NewPostgresUserRepository(db)
	userservice := service.NewUserService(userrepo)
	userhandler := handlers.NewUserHandler(userservice)

	labelrepo := postgres.NewPostgresLabelRepository(db)
	labelservice := service.NewLabelService(labelrepo, activityservice)
	labelhandler := handlers.NewLabelHandler(labelservice)

	commentrepo := postgres.NewPostgresCommentRepository(db)
	commentservice := service.NewCommentService(commentrepo, activityservice)
	commenthandler := handlers.NewCommentHandler(commentservice)

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
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

}

func runMigrations(dbURL string) {
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal("migration init failed:", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("migration failed:", err)
	}
	log.Println("migrations complete")
}
