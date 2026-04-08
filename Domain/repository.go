package domain

type IssueRepository interface {
	NewIssue(issue Issue) error
	GetById(id string) (Issue, error)
	GetByStatus(status []IssueStatus) ([]Issue, error)
	GetByTitle(title string) ([]Issue, error)
	UpdateIssue(issue Issue) error
	DeleteIssue(issue Issue) error
	ListIssues() ([]Issue, error)
}

type LabelRepository interface {
	NewLabel(label Label) error
	GetById(id string) (Label, error)
	GetByTitle(id string) (Label, error)
	GetByColour(colour string) ([]Label, error)
	UpdateLabel(label Label) error
	DeleteLabel(label Label) error
	ListLabels() ([]Label, error)
}

type CommentRepository interface {
	NewComment(comment Comment) error
	GetById(id string) (Comment, error)
	GetByUserId(userid string) ([]Comment, error)
	GetByIssueId(issueid string) ([]Comment, error)
	UpdateComment(comment Comment) error
	DeleteComment(comment Comment) error
	CommentList() ([]Comment, error)
}

type ActivityRepository interface {
	NewActivity(activity Activity) error
	GetByUserId(userid string) ([]Activity, error)
	GetByIssueId(issueid string) ([]Activity, error)
	GetByAction(action ActivityType) ([]Activity, error)
	UpdateActivity(activity Activity) error
	DeleteActivity(activity Activity) error
	ActivityList() ([]Activity, error)
}

type UserRepository interface {
	NewUser(user User) error
	ChangeRole(User Roles) error
	GetByName(name string) (User, error)
	GetById(id string) (User, error)
	GetByRole(role Roles) ([]User, error)
	UpdateRole(user Roles) error
	DeleteUser(user User) error
	UserList() ([]User, error)
}
