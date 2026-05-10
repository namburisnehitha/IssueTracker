package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/namburisnehitha/IssueTracker/domain"
	"github.com/namburisnehitha/IssueTracker/internal/auth"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type UserService struct {
	userRepository domain.UserRepository
	tracer         trace.Tracer
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
		tracer:         otel.Tracer("user-service"),
	}
}

func (u *UserService) CreateUser(ctx context.Context, name string, username string, password string) (string, error) {

	ctx, span := u.tracer.Start(ctx, "CreateUser")
	defer span.End()

	id := uuid.New().String()

	user, err := domain.NewUser(name, id, username, password)
	if err != nil {
		return "", err
	}
	user.Password, err = auth.HashPassword(password)

	if err != nil {
		return "", err
	}

	return id, u.userRepository.Save(ctx, user)
}

func (u *UserService) GetById(ctx context.Context, id string) (domain.User, error) {
	ctx, span := u.tracer.Start(ctx, "GetById")
	defer span.End()
	return u.userRepository.GetById(ctx, id)
}

func (u *UserService) GetByName(ctx context.Context, name string) ([]domain.User, error) {
	ctx, span := u.tracer.Start(ctx, "GetByName")
	defer span.End()
	return u.userRepository.GetByName(ctx, name)
}

func (u *UserService) GetByRole(ctx context.Context, role domain.Roles) ([]domain.User, error) {
	ctx, span := u.tracer.Start(ctx, "GetByRole")
	defer span.End()
	return u.userRepository.GetByRole(ctx, role)
}

func (u *UserService) UpdateUser(ctx context.Context, user domain.User) error {
	ctx, span := u.tracer.Start(ctx, "UpdateUser")
	defer span.End()
	return u.userRepository.UpdateUser(ctx, user)
}

func (u *UserService) DeleteUser(ctx context.Context, user domain.User) error {
	ctx, span := u.tracer.Start(ctx, "DeleteUser")
	defer span.End()
	return u.userRepository.DeleteUser(ctx, user)
}

func (u *UserService) UserList(ctx context.Context) ([]domain.User, error) {
	ctx, span := u.tracer.Start(ctx, "UserList")
	defer span.End()
	return u.userRepository.UserList(ctx)
}

func (u *UserService) GetByUserName(ctx context.Context, username string) (domain.User, error) {
	ctx, span := u.tracer.Start(ctx, "GetByUserName")
	defer span.End()
	return u.userRepository.GetByUserName(ctx, username)
}
