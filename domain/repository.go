package domain

import "context"

type IssueRepository interface {
	Save(ctx context.Context, issue Issue) error
	GetById(ctx context.Context, id string) (Issue, error)
	GetByStatus(ctx context.Context, status IssueStatus) ([]Issue, error)
	GetByTitle(ctx context.Context, title string) ([]Issue, error)
	UpdateIssue(ctx context.Context, issue Issue) error
	DeleteIssue(ctx context.Context, issue Issue) error
	ListIssues(ctx context.Context) ([]Issue, error)
}

type LabelRepository interface {
	Save(ctx context.Context, label Label) error
	GetById(ctx context.Context, id string) (Label, error)
	GetByName(ctx context.Context, id string) (Label, error)
	GetByColour(ctx context.Context, colour string) ([]Label, error)
	UpdateLabel(ctx context.Context, label Label) error
	DeleteLabel(ctx context.Context, label Label) error
	LabelList(ctx context.Context) ([]Label, error)
}

type CommentRepository interface {
	Save(ctx context.Context, comment Comment) error
	GetById(ctx context.Context, id string) (Comment, error)
	GetByUserId(ctx context.Context, userid string) ([]Comment, error)
	GetByIssueId(ctx context.Context, issueid string) ([]Comment, error)
	UpdateComment(ctx context.Context, comment Comment) error
	DeleteComment(ctx context.Context, comment Comment) error
	CommentList(ctx context.Context) ([]Comment, error)
}

type ActivityRepository interface {
	Save(ctx context.Context, activity Activity) error
	GetById(ctx context.Context, id string) (Activity, error)
	GetByUserId(ctx context.Context, userid string) ([]Activity, error)
	GetByIssueId(ctx context.Context, issueid string) ([]Activity, error)
	GetByAction(ctx context.Context, action ActivityType) ([]Activity, error)
	ActivityList(ctx context.Context) ([]Activity, error)
}

type UserRepository interface {
	Save(ctx context.Context, user User) error
	GetByName(ctx context.Context, name string) ([]User, error)
	GetById(ctx context.Context, id string) (User, error)
	GetByRole(ctx context.Context, role Roles) ([]User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, user User) error
	UserList(ctx context.Context) ([]User, error)
	GetByUserName(ctx context.Context, name string) (User, error)
}
