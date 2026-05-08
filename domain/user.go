package domain

import (
	"time"
)

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
	UserName      string
	Password      string
}

func NewUser(name string, id string, username string, password string) (User, error) {
	if name == "" {
		return User{}, ErrInvalidUserData
	}

	return User{
		Name:          name,
		Role:          RoleDeveloper,
		Id:            id,
		JoinedAt:      time.Now(),
		ChangedRoleAt: time.Time{},
		UserName:      username,
		Password:      password,
	}, nil
}

func (u *User) ChangeRole(role Roles) {
	u.Role = role
	u.ChangedRoleAt = time.Now()
}
