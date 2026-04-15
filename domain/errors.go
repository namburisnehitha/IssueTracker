package domain

import (
	"fmt"
)

type DomainError struct {
	Code    string
	Message string
}

func (e DomainError) Error() string {
	return fmt.Sprintf("%s:%s", e.Code, e.Message)
}

var (
	ErrOnlyAssigneeCanStartProgress = DomainError{
		Code:    "ONLY_ASSIGNEE_CAN_START_PROGRESS",
		Message: "Only the assignee can start the progress",
	}
	ErrInvalidStateTransition = DomainError{
		Code:    "INVALID_STATE_TRANSITION",
		Message: "invalid issue state transition",
	}
	ErrIssueAlreadyClosed = DomainError{
		Code:    "ISSUE_ALREADY_CLOSED",
		Message: "issue is already closed",
	}
	ErrIssueNotAssigned = DomainError{
		Code:    "ISSUE_NOT_ASSIGNED",
		Message: "issue is not assigned",
	}
	ErrIssueHasNoAssignee = DomainError{
		Code:    "ISSUE_HAS_NO_ASSIGNEE",
		Message: "issue cannot move to IN_PROGRESS without assignee",
	}
	ErrIssueAlreadyAssigned = DomainError{
		Code:    "ISSUE_ALREADY_ASSIGNED",
		Message: "issue is already assigned ",
	}
	ErrUnauthorizedAction = DomainError{
		Code:    "UNAUTHORISED_ACTION",
		Message: "only certain users can perform certain tasks",
	}
	ErrIssueNotFound = DomainError{
		Code:    "ISSUE_NOT_FOUND",
		Message: "operations requires a valid issue",
	}
	ErrInvalidIssueData = DomainError{
		Code:    "INVALID_ISSUE_DATA",
		Message: "issue data is not valid",
	}
	ErrInvalidLabelData = DomainError{
		Code:    "INVALID_LABEL_DATA",
		Message: "label data is not valid",
	}
	ErrInvalidCommentData = DomainError{
		Code:    "INVALID_COMMENT_DATA",
		Message: "comment data is not valid",
	}
	ErrInvalidActivityData = DomainError{
		Code:    "INVALID_ACTIVITY_DATA",
		Message: "activity data is not valid",
	}
	ErrInvalidUserData = DomainError{
		Code:    "INVALID_USER_DATA",
		Message: "user data is not valid",
	}
)
