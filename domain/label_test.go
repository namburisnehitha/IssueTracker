package domain

import (
	"testing"
)

func TestNewLabel(t *testing.T) {
	//Invalid data
	id2 := "01"
	name2 := ""
	description2 := "no Title"
	colour2 := "colour"
	_, err := NewLabel(id2, name2, description2, colour2)

	if err != ErrInvalidLabelData {
		t.Errorf("got %v,want %v", ErrInvalidLabelData, err)
	}

	id := "01"
	name := "bug"
	description := "need fix"
	colour := "red"

	label, err := NewLabel(id, name, description, colour)

	if err != nil {
		t.Fatalf("failed to create label: %v", err)
	}

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

	label, err := NewLabel(id, name, description, colour)

	if err != nil {
		t.Fatalf("failed to create label: %v", err)
	}

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

	label, err := NewLabel(id, name, description, colour)

	if err != nil {
		t.Fatalf("failed to create label: %v", err)
	}

	label.UpdateColour(new_colour)

	if label.Colour != new_colour {
		t.Errorf("got %v,want %v", label.Name, new_colour)
	}

}
