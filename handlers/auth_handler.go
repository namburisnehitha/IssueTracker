package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/namburisnehitha/IssueTracker/internal/auth"
	"github.com/namburisnehitha/IssueTracker/service"
)

type AuthHandler struct {
	userService *service.UserService
}

type CreateUserLoginInfoRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(userservice *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userservice,
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var ur *CreateUserLoginInfoRequest
	err := json.NewDecoder(r.Body).Decode(&ur)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ah.userService.GetByUserName(ur.UserName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ok := auth.CheckPasswordHash(ur.Password, user.Password)
	if !ok {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	writeJSON(w, http.StatusOK, token)

}
