package service

import "github.com/namburisnehitha/IssueTracker/domain"

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) CreateUser(id string, name string, role domain.Roles) error {
	user, err := domain.NewUser(name, role, id)
	if err != nil {
		return err
	}
	return u.userRepository.Save(user)
}

func (u *UserService) GetById(id string) (domain.User, error) {
	return u.userRepository.GetById(id)
}

func (u *UserService) GetByName(name string) ([]domain.User, error) {
	return u.userRepository.GetByName(name)
}

func (u *UserService) GetByRole(role domain.Roles) ([]domain.User, error) {
	return u.userRepository.GetByRole(role)
}

func (u *UserService) UpdateUser(user domain.User) error {
	return u.userRepository.UpdateUser(user)
}

func (u *UserService) DeleteUser(user domain.User) error {
	return u.userRepository.DeleteUser(user)
}

func (u *UserService) UserList() ([]domain.User, error) {
	return u.userRepository.UserList()
}
