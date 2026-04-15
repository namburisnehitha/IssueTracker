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
		t.Errorf("got %v,want %v", err, ErrInvalidUserData)
	}

	name := "snehitha"
	role := RoleDeveloper
	id := "01"
	user, err := NewUser(name, role, id)

	if user.Name != name {
		t.Errorf("got %v,want %v", user.Name, name)
	}
	if user.Role != role {
		t.Errorf("got %v,want %v", user.Role, role)
	}
	if user.Id != id {
		t.Errorf("got %v,want %v", user.Id, id)
	}

	if user.JoinedAt.IsZero() {
		t.Errorf("got %v,want %v", user.JoinedAt, time.Now())
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestChangeRole(t *testing.T) {
	name := "snehitha"
	role := RoleDeveloper
	id := "01"
	new_role := RoleMaintainer
	user, err := NewUser(name, role, id)
	user.ChangeRole(new_role)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
	if user.Role != new_role {
		t.Errorf("got %v, want %v", user.Role, new_role)
	}

}
