package domain

import "time"

type Roles string

const (
	RoleAdmin      Roles = "ADMIN"
	RoleMaintainer Roles = "MAINTAINER"
	RoleDeveloper  Roles = "DEVELOPER"
)

type User struct {
	Name     string
	Role     Roles
	Id       string
	JoinedAt time.Time
}

func (u *User) ChangeRole(role Roles) {
	u.Role = role
}
