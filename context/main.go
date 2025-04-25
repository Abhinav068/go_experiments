package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create parent context
	parent := context.Background()

	// Create cancellable context
	ctx, cancel := context.WithCancel(parent)

	// Start multiple operations
	go DoWork(ctx, "Task 1")
	go DoWork(ctx, "Task 2")

	time.Sleep(2 * time.Second)
	// Cancel all operations
	cancel()
	time.Sleep(1 * time.Second) // Wait to see effect
}

func DoWork(ctx context.Context, task string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s cancelled\n", task)
			return
		default:
			fmt.Printf("Working on %s\n", task)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
