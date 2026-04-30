package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/namburisnehitha/IssueTracker/service"
)

type CreateLabelRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Colour      string `json:"colour"`
}

type LabelHandler struct {
	labelService service.LabelService
}

func NewLabelHandler(labelSevice *service.LabelService) *LabelHandler {
	return &LabelHandler{
		labelService: *labelSevice,
	}
}

func (l *LabelHandler) CreateLabel(w http.ResponseWriter, r *http.Request) {

	var lr CreateLabelRequest
	err := json.NewDecoder(r.Body).Decode(&lr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lr.Id, err = l.labelService.CreateLabel(lr.Name, lr.Description, lr.Colour)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, nil)
}

func (l *LabelHandler) GetById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	label, err := l.labelService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, label)
}

func (l *LabelHandler) GetByName(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	label, err := l.labelService.GetByName(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, label)
}

func (l *LabelHandler) GetByColour(w http.ResponseWriter, r *http.Request) {

	colour := r.URL.Query().Get("colour")
	label, err := l.labelService.GetByColour(colour)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, label)
}

func (l *LabelHandler) UpdateLabel(w http.ResponseWriter, r *http.Request) {

	var lr CreateLabelRequest
	err := json.NewDecoder(r.Body).Decode(&lr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	label, err := l.labelService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	label.Colour = lr.Colour
	label.Name = lr.Name
	label.Description = lr.Description
	err = l.labelService.UpdateLabel(label)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, label)
}

func (l *LabelHandler) DeleteLabel(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	label, err := l.labelService.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = l.labelService.DeleteLabel(label)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (l *LabelHandler) LabelList(w http.ResponseWriter, r *http.Request) {

	labels, err := l.labelService.LabelList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, labels)
}
