package main

import (
	"fmt"
	"github.com/namburisnehitha/IssueTracker/domain"
	"github.com/namburisnehitha/IssueTracker/internal/postgres"
)

func main() {
	db, err := postgres.NewDB("postgres://postgres:darealsql@localhost:5432/issuetracker?sslmode=disable")
	if err != nil {
		return
	}
	userrepo := postgres.NewPostgresUserRepository(db)
	user, err := domain.NewUser("sneh", "id")
	if err != nil {
		return
	}

	err = userrepo.Save(user)
	if err != nil {
		fmt.Println("Save error:", err)
	} else {
		fmt.Println("Saved successfully")
	}
}
