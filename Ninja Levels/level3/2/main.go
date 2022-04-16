package main

import (
	"context"
	"fmt"

	"github.com/AlexL70/golang-architecture/tree/main/Ninja_Levels/level3/2/session"
)

func main() {
	ctx := context.Background()
	ctx = session.SetUserID(ctx, 12345)
	ctx = session.SetAdminAccess(ctx, true)
	userID := session.GetUserID(ctx)
	isAdmin := session.GetAdmin(ctx)
	fmt.Println("User name is", userID)
	fmt.Println("Is he admin? It is", isAdmin)
}
