package domain

type IssueRepository interface {
	Save(issue Issue) error
	GetById(id string) (Issue, error)
	GetByStatus(status IssueStatus) ([]Issue, error)
	GetByTitle(title string) ([]Issue, error)
	UpdateIssue(issue Issue) error
	DeleteIssue(issue Issue) error
	ListIssues() ([]Issue, error)
}

type LabelRepository interface {
	Save(label Label) error
	GetById(id string) (Label, error)
	GetByName(id string) (Label, error)
	GetByColour(colour string) ([]Label, error)
	UpdateLabel(label Label) error
	DeleteLabel(label Label) error
	LabelList() ([]Label, error)
}

type CommentRepository interface {
	Save(comment Comment) error
	GetById(id string) (Comment, error)
	GetByUserId(userid string) ([]Comment, error)
	GetByIssueId(issueid string) ([]Comment, error)
	UpdateComment(comment Comment) error
	DeleteComment(comment Comment) error
	CommentList() ([]Comment, error)
}

type ActivityRepository interface {
	Save(activity Activity) error
	GetById(id string) (Activity, error)
	GetByUserId(userid string) ([]Activity, error)
	GetByIssueId(issueid string) ([]Activity, error)
	GetByAction(action ActivityType) ([]Activity, error)
	ActivityList() ([]Activity, error)
}

type UserRepository interface {
	Save(user User) error
	GetByName(name string) ([]User, error)
	GetById(id string) (User, error)
	GetByRole(role Roles) ([]User, error)
	UpdateUser(user User) error
	DeleteUser(user User) error
	UserList() ([]User, error)
}
