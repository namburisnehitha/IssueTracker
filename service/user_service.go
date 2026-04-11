package service

import "github.com/namburisnehitha/IssueTracker/domain"

type UserService struct {
	UserRepository domain.UserRepository
}

func NewUserService(UserRepository domain.UserRepository) UserService {
	return UserService{
		UserRepository: UserRepository,
	}
}
