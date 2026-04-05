
func TestNewIssue(t *testing.T) {
	id := "ISS-001"
	title := "Login broken"
	description := "Users cannot log in"

	issue := NewIssue(id, title, description)

	if issue.Id != id {
		t.Errorf("got %v,want %v", issue.Id, id)
	}

	if issue.Title != title {
		t.Errorf("got %v,want %v", issue.Title, title)
	}

	if issue.Description != description {
		t.Errorf("got %v,want %v", issue.Description, description)
	}

	if issue.Status != StatusOpen {
		t.Errorf("got %v,want %v", issue.Status, StatusOpen)
	}

	if issue.CreatedAt.IsZero() {
		t.Errorf("got %v,want %v", issue.CreatedAt, time.Now())
	}
}