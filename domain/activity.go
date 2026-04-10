package domain

import "time"

type ActivityType string

const (
	UserAssigned  ActivityType = "USER_ASSIGNED"
	StatusChanged ActivityType = "STATUS_CHANGED"
	CommentAdded  ActivityType = "COMMENT_ADDED"
	LabelAdded    ActivityType = "LABEL_ADDED"
)

type Activity struct {
	UserId      string
	IssueId     string
	Description string
	CreatedAt   time.Time
	Action      ActivityType
}

func NewActivity(IssueId string, UserId string, Description string, Action ActivityType) Activity {
	return Activity{
		IssueId:     IssueId,
		UserId:      UserId,
		Description: Description,
		Action:      Action,
		CreatedAt:   time.Now(),
	}
}
