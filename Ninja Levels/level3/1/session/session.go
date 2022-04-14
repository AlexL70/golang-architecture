package session

import "context"

type key string

const userIdKey key = "userIdKey"
const userNameKey key = "retuserNameKey"

func WithUserID(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, userIdKey, userId)
}

func WithUserName(ctx context.Context, userName string) context.Context {
	return context.WithValue(ctx, userNameKey, userName)
}

func GetUserId(ctx context.Context) *int {
	userId, ok := ctx.Value(userIdKey).(int)
	if !ok {
		return nil
	}
	return &userId
}

func GetUserName(ctx context.Context) *string {
	userName, ok := ctx.Value(userNameKey).(string)
	if !ok {
		return nil
	}
	return &userName
}
