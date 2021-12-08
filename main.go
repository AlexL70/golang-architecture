package main

import (
	"context"
	"fmt"
)

type key string

var userKey key = "userKey"

func main() {
	ctx := context.WithValue(context.Background(), userKey, 1)
	//ctx := context.Background()

	userId, ok := ctx.Value(userKey).(int)
	if !ok {
		fmt.Print("Not logged in")
		return
	}
	fmt.Print("UserId = ", userId)
}
