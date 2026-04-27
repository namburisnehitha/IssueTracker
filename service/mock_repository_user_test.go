package service

import (
	"testing"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type MockUserRepository struct {
	users map[string]domain.User
}

func (m *MockUserRepository) Save(user domain.User) error {
	m.users[user.Id] = user
	return nil
}

func (m *MockUserRepository) GetById(id string) (domain.User, error) {
	user := m.users[id]
	return user, nil
}

func (m *MockUserRepository) GetByName(name string) ([]domain.User, error) {
	var result []domain.User
	for _, user := range m.users {
		if user.Name == name {
			result = append(result, user)
		}
	}
	return result, nil
}

func (m *MockUserRepository) GetByRole(role domain.Roles) ([]domain.User, error) {
	var result []domain.User
	for _, user := range m.users {
		if user.Role == role {
			result = append(result, user)
		}
	}
	return result, nil
}

func (m *MockUserRepository) UpdateUser(user domain.User) error {
	m.users[user.Id] = user
	return nil
}

func (m *MockUserRepository) DeleteUser(user domain.User) error {
	delete(m.users, user.Id)
	return nil
}

func (m *MockUserRepository) UserList() ([]domain.User, error) {
	var result []domain.User
	for _, user := range m.users {
		result = append(result, user)
	}
	return result, nil
}

func TestCreateUser(t *testing.T) {

	name := "user"
	repo := &MockUserRepository{users: map[string]domain.User{}}
	service := NewUserService(repo)
	id, err := service.CreateUser(name)
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
	user, err := service.GetById(id)

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
	users, err := service.GetByName(name)

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
	users, err := service.GetByRole(role)

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
	err := service.UpdateUser(user)
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
	err := service.DeleteUser(user)

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
	users, err := service1.UserList()

	if len(users) != 3 {
		t.Errorf("got %v, want %v", len(users), 3)
	}
	if err != nil {
		t.Errorf("got %v,want %v", err, nil)
	}
}
