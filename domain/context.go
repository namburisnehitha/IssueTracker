package domain

import "context"

type contextKey string

const UserIDKey contextKey = "userID"

func UserIDFromContext(ctx context.Context) string {
	id, _ := ctx.Value(UserIDKey).(string)
	return id
}
