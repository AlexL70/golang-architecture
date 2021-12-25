package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancelF := context.WithCancel(ctx)

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println("running", n)
			for {
				select {
				case <-ctx.Done():
					fmt.Println("close", n)
					return
				default:
					fmt.Println("still working", n)
					time.Sleep(100 * time.Millisecond)
				}
				fmt.Println("GOROUTINS RUNNING:", runtime.NumGoroutine())
			}
		}(i)
	}
	time.Sleep(3 * time.Second)
	cancelF()
	fmt.Println("Sleeping for 1 seconds")
	time.Sleep(time.Second)
	fmt.Println("GOROUTINS RUNNING:", runtime.NumGoroutine())
}
