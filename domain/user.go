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

func NewUser(Name string, role Roles, id string) (User, error) {
	if Name == "" {
		return User{}, ErrInvalidUserData
	}
	return User{
		Name:          Name,
		Role:          role,
		Id:            id,
		JoinedAt:      time.Now(),
		ChangedRoleAt: time.Time{},
	}, nil
}

func (u *User) ChangeRole(role Roles) {
	u.Role = role
	u.ChangedRoleAt = time.Now()
}
