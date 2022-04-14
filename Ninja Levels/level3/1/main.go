package main

import (
	"context"
	"fmt"

	"github.com/AlexL70/golang-architecture/tree/main/Ninja_Levels/level3/1/session"
)

func main() {
	ctx := context.Background()
	ctx = session.WithUserID(ctx, 12)
	userId := session.GetUserId(ctx)
	if userId == nil {
		fmt.Println("Not logget in")
		return
	}
	fmt.Println("User ID = ", *userId)
	ctx = session.WithUserName(ctx, "AlexL")
	userName := session.GetUserName(ctx)
	if userName == nil {
		fmt.Println("User name is unknown")
		return
	}
	fmt.Println("User name is ", *userName)
}
