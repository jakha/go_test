package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	sleepAndTalk(ctx, 1*time.Second, "solo")
	cancel()
}

func sleepAndTalk(ctx context.Context, t time.Duration, msg string) {
	select {
	case <-time.After(t):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
