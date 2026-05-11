package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/namburisnehitha/IssueTracker/domain"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func domainErrorToStatus(err error) int {
	var de domain.DomainError
	if !errors.As(err, &de) {
		return http.StatusInternalServerError
	}
	switch de.Code {
	case "ISSUE_NOT_FOUND":
		return http.StatusNotFound
	case "ISSUE_ALREADY_ASSIGNED", "ISSUE_ALREADY_CLOSED":
		return http.StatusConflict
	case "INVALID_ISSUE_DATA", "INVALID_LABEL_DATA",
		"INVALID_COMMENT_DATA", "INVALID_ACTIVITY_DATA",
		"INVALID_USER_DATA":
		return http.StatusUnprocessableEntity
	case "UNAUTHORISED_ACTION":
		return http.StatusForbidden
	case "INVALID_STATE_TRANSITION":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
