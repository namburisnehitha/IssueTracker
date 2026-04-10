package domain

import (
	"testing"
)

func TestChangeRole(t *testing.T) {
	name := "snehitha"
	role := RoleDeveloper
	id := "01"
	new_role := RoleMaintainer
	user := User{Name: name, Role: role, Id: id}
	user.ChangeRole(new_role)
	if user.Role != new_role {
		t.Errorf("got %v, want %v", user.Role, new_role)
	}
}
