package domain

import (
	"testing"
)

func TestNewLabel(t *testing.T) {
	id := "01"
	name := "bug"
	description := "need fix"
	colour := "red"

	label := NewLabel(id, name, description, colour)

	if label.Id != id {
		t.Errorf("got %v,want %v", label.Id, id)
	}

	if label.Name != name {
		t.Errorf("got %v,want %v", label.Name, name)
	}

	if label.Description != description {
		t.Errorf("got %v,want %v", label.Description, description)
	}

	if label.Colour != colour {
		t.Errorf("got %v,want %v", label.Colour, colour)
	}
}

func TestUpdateName(t *testing.T) {
	id := "01"
	name := "bug"
	description := "need fix"
	colour := "red"
	new_name := "urgent"

	label := NewLabel(id, name, description, colour)
	label.UpdateName(new_name)
	if label.Name != new_name {
		t.Errorf("got %v,want %v", label.Name, new_name)
	}
}

func TestUpdateColour(t *testing.T) {
	id := "01"
	name := "bug"
	description := "need fix"
	colour := "red"
	new_colour := "pink"

	label := NewLabel(id, name, description, colour)
	label.UpdateColour(new_colour)
	if label.Colour != new_colour {
		t.Errorf("got %v,want %v", label.Name, new_colour)
	}
}
