package service

import (
	"context"
	"testing"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type MockUserRepository struct {
	users map[string]domain.User
}

func (m *MockUserRepository) Save(ctx context.Context, user domain.User) error {
	m.users[user.Id] = user
	return nil
}

func (m *MockUserRepository) GetById(ctx context.Context, id string) (domain.User, error) {
	user := m.users[id]
	return user, nil
}

func (m *MockUserRepository) GetByName(ctx context.Context, name string) ([]domain.User, error) {
	var result []domain.User
	for _, user := range m.users {
		if user.Name == name {
			result = append(result, user)
		}
	}
	return result, nil
}

func (m *MockUserRepository) GetByRole(ctx context.Context, role domain.Roles) ([]domain.User, error) {
	var result []domain.User
	for _, user := range m.users {
		if user.Role == role {
			result = append(result, user)
		}
	}
	return result, nil
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user domain.User) error {
	m.users[user.Id] = user
	return nil
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, user domain.User) error {
	delete(m.users, user.Id)
	return nil
}

func (m *MockUserRepository) UserList(ctx context.Context) ([]domain.User, error) {
	var result []domain.User
	for _, user := range m.users {
		result = append(result, user)
	}
	return result, nil
}

func (m *MockUserRepository) GetByUserName(ctx context.Context, username string) (domain.User, error) {
	for _, user := range m.users {
		if user.UserName == username {
			return user, nil
		}
	}
	return domain.User{}, domain.ErrInvalidUserData
}

func TestCreateUser(t *testing.T) {

	name := "user"
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service := NewUserService(repo)
	id, err := service.CreateUser(context.Background(), name, "username", "##")
	saved := repo.users[id]

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if saved.Id != id {
		t.Errorf("got %v,want %v", saved.Id, id)
	}

	if saved.Name != name {
		t.Errorf("got %v,want %v", saved.Name, name)
	}

}

func TestGetById(t *testing.T) {
	id := "01"
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service := NewUserService(repo)
	repo.users[id] = domain.User{Id: id, Name: "user", Role: domain.RoleAdmin}
	user, err := service.GetById(context.Background(), id)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if user.Id != id {
		t.Errorf("got %v,want %v", user.Id, id)
	}
}
func TestGetByName(t *testing.T) {
	name := "user1"
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service := NewUserService(repo)
	repo.users["01"] = domain.User{Id: "01", Name: "user1"}
	repo.users["02"] = domain.User{Id: "02", Name: "user2"}
	repo.users["03"] = domain.User{Id: "03", Name: "user1"}
	users, err := service.GetByName(context.Background(), name)

	for _, user := range users {
		if user.Name != name {
			t.Errorf("got %v,want %v", user.Name, name)
		}
	}

	if len(users) != 2 {
		t.Errorf("got %v,want %v", len(users), 2)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}
func TestGetByRole(t *testing.T) {
	role := domain.RoleAdmin
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service := NewUserService(repo)
	repo.users["01"] = domain.User{Id: "01", Name: "user1", Role: domain.RoleAdmin}
	repo.users["02"] = domain.User{Id: "02", Name: "user2", Role: domain.RoleDeveloper}
	repo.users["03"] = domain.User{Id: "03", Name: "user3", Role: domain.RoleAdmin}
	users, err := service.GetByRole(context.Background(), role)

	for _, user := range users {
		if user.Role != role {
			t.Errorf("got %v,want %v", user.Role, role)
		}
	}

	if len(users) != 2 {
		t.Errorf("got %v,want %v", len(users), 2)
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestUpdateUser(t *testing.T) {
	newrole := domain.RoleAdmin
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service := NewUserService(repo)
	repo.users["01"] = domain.User{Id: "01", Name: "user", Role: domain.RoleDeveloper}
	user := domain.User{Id: "01", Name: "user", Role: newrole}
	err := service.UpdateUser(context.Background(), user)
	saved := repo.users["01"]

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if saved.Role != newrole {
		t.Errorf("got %v,want %v", saved.Role, newrole)
	}

}

func TestDeleteUser(t *testing.T) {
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service := NewUserService(repo)
	repo.users["01"] = domain.User{Id: "01", Name: "user", Role: domain.RoleDeveloper}
	user := domain.User{Id: "01"}
	err := service.DeleteUser(context.Background(), user)

	_, exists := repo.users["01"]

	if exists {
		t.Errorf("issue was not deleted")
	}

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}
func TestUserList(t *testing.T) {
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service1 := NewUserService(repo)
	repo.users["01"] = domain.User{Id: "01", Name: "name"}
	repo.users["02"] = domain.User{Id: "02", Name: "name"}
	repo.users["03"] = domain.User{Id: "03", Name: "name"}
	users, err := service1.UserList(context.Background())

	if len(users) != 3 {
		t.Errorf("got %v, want %v", len(users), 3)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}

func TestGetByUserName(t *testing.T) {
	id := "01"
	username := "sneh"
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service := NewUserService(repo)
	repo.users[id] = domain.User{Id: id, Name: "user", Role: domain.RoleAdmin, UserName: username}

	user, err := service.GetByUserName(context.Background(), username)

	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}

	if user.UserName != username {
		t.Errorf("got %v,want %v", user.UserName, username)
	}
}
