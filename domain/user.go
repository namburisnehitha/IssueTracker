package domain

import "time"

type Roles string

const (
	RoleAdmin      Roles = "ADMIN"
	RoleMaintainer Roles = "MAINTAINER"
	RoleDeveloper  Roles = "DEVELOPER"
)

type User struct {
	Name          string
	Role          Roles
	Id            string
	JoinedAt      time.Time
	ChangedRoleAt time.Time
}

func NewUser(name string, role Roles, id string) (User, error) {
	if name == "" {
		return User{}, ErrInvalidUserData
	}
	return User{
		Name:          name,
		Role:          role,
		Id:            id,
		JoinedAt:      time.Now(),
		ChangedRoleAt: time.Now(),
	}, nil
}

func (u *User) ChangeRole(role Roles) {
	u.Role = role
}
