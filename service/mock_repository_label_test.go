package service

import (
	"testing"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type MockLabelRepository struct {
	labels map[string]domain.Label
}

func (m *MockLabelRepository) Save(label domain.Label) error {
	m.labels[label.Id] = label
	return nil
}

func (m *MockLabelRepository) GetById(id string) (domain.Label, error) {
	label := m.labels[id]
	return label, nil
}

func (m *MockLabelRepository) GetByName(name string) (domain.Label, error) {
	label := m.labels[name]
	return label, nil
}

func (m *MockLabelRepository) GetByColour(colour string) ([]domain.Label, error) {
	var result []domain.Label
	for _, label := range m.labels {
		if label.Colour == colour {
			result = append(result, label)
		}
	}
	return result, nil
}

func (m *MockLabelRepository) UpdateLabel(label domain.Label) error {
	m.labels[label.Id] = label
	return nil
}

func (m *MockLabelRepository) DeleteLabel(label domain.Label) error {
	delete(m.labels, label.Id)
	return nil
}

func (m *MockLabelRepository) LabelList() ([]domain.Label, error) {
	var result []domain.Label
	for _, label := range m.labels {
		result = append(result, label)
	}
	return result, nil
}
func TestCreateLabel(t *testing.T) {
	id := "01"
	name := "name"
	description := "description"
	colour := "colour"
	repo := &MockLabelRepository{labels: map[string]domain.Label{}}
	service := NewLabelService(repo)
	repo.labels[id] = domain.Label{Id: id, Name: name, Description: description, Colour: colour}
	err := service.CreateLabel(id, name, description, colour)
	saved := repo.labels[id]

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if saved.Id != id {
		t.Errorf("got %v,want %v", saved.Id, id)
	}

	if saved.Name != name {
		t.Errorf("got %v,want %v", saved.Name, name)
	}

	if saved.Description != description {
		t.Errorf("got %v,want %v", saved.Description, description)
	}

	if saved.Colour != colour {
		t.Errorf("got %v,want %v", saved.Colour, colour)
	}
}

func TestLabelGetById(t *testing.T) {
	id := "01"
	repo := &MockLabelRepository{labels: map[string]domain.Label{}}
	service := NewLabelService(repo)
	repo.labels["01"] = domain.Label{Id: id}
	label, err := service.GetById(id)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
	if label.Id != id {
		t.Errorf("got %v,want %v", label.Id, id)
	}
}

func TestLabelGetByName(t *testing.T) {
	id := "01"
	name := "title"
	repo := &MockLabelRepository{labels: map[string]domain.Label{}}
	service := NewLabelService(repo)
	repo.labels[name] = domain.Label{Id: id, Name: name}
	label, err := service.GetByName(name)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
	if label.Name != name {
		t.Errorf("got %v,want %v", label.Id, id)
	}
}
func TestLabelGetByColour(t *testing.T) {
	colour := "red"
	repo := &MockLabelRepository{labels: map[string]domain.Label{}}
	service := NewLabelService(repo)
	var labels []domain.Label
	repo.labels["01"] = domain.Label{Id: "01", Colour: "red"}
	repo.labels["02"] = domain.Label{Id: "02", Colour: "pink"}
	repo.labels["03"] = domain.Label{Id: "03", Colour: "red"}
	labels, err := service.GetByColour(colour)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	for _, label := range labels {
		if label.Colour != colour {
			t.Errorf("got %v,want %v", label.Colour, colour)
		}
	}
	if len(labels) != 2 {
		t.Errorf("got %v,want %v", len(labels), 2)
	}

}
func TestUpdatelabel(t *testing.T) {
	colour := "newcolour"
	repo := &MockLabelRepository{labels: map[string]domain.Label{}}
	service := NewLabelService(repo)
	repo.labels["01"] = domain.Label{Id: "01", Colour: "oldcolour"}
	label := domain.Label{Id: "01", Colour: colour}
	err := service.UpdateLabel(label)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if label.Colour != colour {
		t.Errorf("got %v,want %v", label.Colour, colour)
	}

}

func TestDeleteLabel(t *testing.T) {

	id := "01"
	repo := &MockLabelRepository{labels: map[string]domain.Label{}}
	service := NewLabelService(repo)
	repo.labels[id] = domain.Label{Id: id}
	label := domain.Label{Id: id}
	err := service.DeleteLabel(label)

	_, exists := repo.labels["01"]

	if exists {
		t.Errorf("issue was not deleted")
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestListLabels(t *testing.T) {
	repo := &MockLabelRepository{labels: map[string]domain.Label{}}
	service := NewLabelService(repo)
	var labels []domain.Label
	repo.labels["01"] = domain.Label{Id: "01", Colour: "red"}
	repo.labels["02"] = domain.Label{Id: "02", Colour: "pink"}
	repo.labels["03"] = domain.Label{Id: "03", Colour: "red"}
	labels, err := service.LabelList()

	if len(labels) != 3 {
		t.Errorf("got %v,want %v", len(labels), 3)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

}
