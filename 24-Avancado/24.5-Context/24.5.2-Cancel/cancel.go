package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id uint8) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("encerrando worker %d...\n", id)
			return
		default:
			fmt.Printf("continuando worker %d...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(2 * time.Second)

	cancel()
	time.Sleep(time.Second)
}
