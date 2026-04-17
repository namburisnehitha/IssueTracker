package domain

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {

	name1 := ""
	role1 := RoleDeveloper
	id1 := "01"
	_, err := NewUser(name1, role1, id1)

	if err != ErrInvalidUserData {
		t.Errorf("got %v,want %v", ErrInvalidUserData, err)
	}

	name := "snehitha"
	role := RoleDeveloper
	id := "01"
	user, err := NewUser(name, role, id)

	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	if user.Name != name {
		t.Errorf("got %v,want %v", user.Name, name)
	}
	if user.Role != role {
		t.Errorf("got %v,want %v", user.Role, role)
	}
	if user.Id != id {
		t.Errorf("got %v,want %v", user.Id, id)
	}

}

func TestChangeRole(t *testing.T) {
	name := "snehitha"
	role := RoleDeveloper
	id := "01"
	new_role := RoleMaintainer
	user, err := NewUser(name, role, id)

	if err != nil {
		t.Fatalf("failed to create a user %v", err)
	}

	before := time.Now()
	user.ChangeRole(new_role)
	after := time.Now()

	if user.Role != new_role {
		t.Errorf("got %v, want %v", user.Role, new_role)
	}

	if user.ChangedRoleAt.Before(before) || user.ChangedRoleAt.After(after) {
		t.Errorf("ChangedRoleAt not set correctly")
	}

}
