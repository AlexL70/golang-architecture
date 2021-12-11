package session

import "context"

type key string

const userKey key = "userKey"

func WithUserID(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, userKey, userId)
}

func GetUserId(ctx context.Context) *int {
	userID, ok := ctx.Value(userKey).(int)
	if !ok {
		return nil
	}
	return &userID
}
