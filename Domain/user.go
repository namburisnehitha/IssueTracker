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

func (u User) UserInfo() string {
	return u.Name + "  " + string(u.Role)
}

func (u *User) ChangeRole(roles Roles) {
	u.Role = roles
}
