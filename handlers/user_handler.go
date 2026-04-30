package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/domain"
	"github.com/namburisnehitha/IssueTracker/service"
)

type CreateUserRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserHandler struct {
	userservice *service.UserService
}

func NewUserHandler(userservice *service.UserService) *UserHandler {
	return &UserHandler{
		userservice: userservice,
	}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var ur CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&ur)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ur.Id, err = u.userservice.CreateUser(ur.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, ur.Id)
}

func (u *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	user, err := u.userservice.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (u *UserHandler) GetByName(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	user, err := u.userservice.GetByName(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (u *UserHandler) GetByRole(w http.ResponseWriter, r *http.Request) {

	role := r.URL.Query().Get("role")
	user, err := u.userservice.GetByRole(domain.Roles(role))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var ur CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&ur)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := chi.URLParam(r, "id")
	user, err := u.userservice.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Name = ur.Name

	err = u.userservice.UpdateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, user)

}

func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	user, err := u.userservice.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.userservice.DeleteUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (u *UserHandler) UserList(w http.ResponseWriter, r *http.Request) {

	users, err := u.userservice.UserList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, users)
}
