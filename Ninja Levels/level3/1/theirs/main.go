package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 12345)
	ctx = context.WithValue(ctx, 1, "admin")
	v := ctx.Value("userID")
	if v == nil {
		fmt.Println("no value is associated with this key")
	} else {
		fmt.Println(v)
	}
	v = ctx.Value(1)
	if v == nil {
		v = ctx.Value(1)
		fmt.Println("no value is associated with this key")
	} else {
		fmt.Println(v)
	}
	v = ctx.Value(2)
	if v == nil {
		fmt.Println("no value is associated with this key")
	} else {
		fmt.Println(v)
	}
}
