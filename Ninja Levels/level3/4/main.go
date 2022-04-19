package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	timeout := 1000
	if len(os.Args) > 1 {
		var err error
		timeout, err = strconv.Atoi(os.Args[1])
		if err != nil {
			timeout = 1000
			fmt.Printf("Wrong timeout value passes: %s. Setting to %d\n", os.Args[1], timeout)
		}
	}
	fmt.Printf("Timeout = %dms\n", timeout)
	ctx := context.Background()
	ctx, cancelF := context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
	defer cancelF()

	slowOp(ctx)
}

func slowOp(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("Slow operation has been cancelled.")
			return
		default:
			fmt.Printf("Slow operation is continued... i = %d\n", i)
		}
	}
	fmt.Println("Hello from slow operation!")
}
