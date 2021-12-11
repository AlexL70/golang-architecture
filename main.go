package main

import (
	"context"
	"fmt"

	"github.com/AlexL70/golang-architecture/session"
)

func main() {
	ctx := context.Background()
	ctx = session.WithUserID(ctx, 1)

	userId := session.GetUserId(ctx)
	if userId == nil {
		fmt.Print("Not logged in")
		return
	}
	fmt.Print("UserId = ", *userId)
}
