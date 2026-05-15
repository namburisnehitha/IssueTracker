package domain

import "time"

type ActivityType string

const (
	UserAssigned   ActivityType = "USER_ASSIGNED"
	StatusChanged  ActivityType = "STATUS_CHANGED"
	CommentAdded   ActivityType = "COMMENT_ADDED"
	LabelAdded     ActivityType = "LABEL_ADDED"
	IssueCreated   ActivityType = "ISSUE_CREATED"
	IssueDeleted   ActivityType = "ISSUE_DELETED"
	LabelRemoved   ActivityType = "LABEL_REMOVED"
	CommentDeleted ActivityType = "COMMENT_DELETED"
)

type Activity struct {
	Id          string
	UserId      string
	IssueId     string
	Description string
	CreatedAt   time.Time
	Action      ActivityType
}

func NewActivity(Id string, IssueId string, UserId string, Description string, Action ActivityType) (Activity, error) {

	if Description == "" {
		return Activity{}, ErrInvalidActivityData
	}
	return Activity{
		Id:          Id,
		IssueId:     IssueId,
		UserId:      UserId,
		Description: Description,
		Action:      Action,
		CreatedAt:   time.Now(),
	}, nil
}
