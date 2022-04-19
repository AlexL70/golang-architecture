package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("%d go routines are currently running\n", runtime.NumGoroutine())
	ctx := context.Background()
	var cancelF context.CancelFunc
	ctx, cancelF = context.WithCancel(ctx)
	defer cancelF()
	for i := 0; i < 100; i++ {
		go func(index int) {
			fmt.Println("launching go routine", index)
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("go routine #%d has finished working\n", index)
					runtime.Goexit()
				default:
					// do some work
					time.Sleep(100 * time.Millisecond)
					fmt.Printf("go routine #%d has done some work\n", index)
				}
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Printf("%d go routines are currently running\n", runtime.NumGoroutine())
	time.Sleep(300 * time.Millisecond)
	cancelF()
	fmt.Println("Context has been cancelled")
	for {
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("%d go routines are currently running\n", runtime.NumGoroutine())
		if runtime.NumGoroutine() == 1 {
			return
		}
	}
}
