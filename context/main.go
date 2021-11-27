package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	contextBackground()
	contextWithTimeOut()
	contextWithDeadline()
}

func contextBackground() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	mySleepAndTalk(ctx, time.Second*5, "Hello from contextBackground")
}

func contextWithTimeOut() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel() //context package would allocate some resources for timeout, so it's important to clean it up
	mySleepAndTalk(ctx, time.Second*5, "Hello from contextWithTimeOut")
}

func contextWithDeadline() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Hour))
	defer cancel() //context package would allocate some resources for timeout, so it's important to clean it up
	mySleepAndTalk(ctx, time.Second*5, "Hello from contextWithDeadline()")
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	// time.After returns an event in the channel after specified amount of time
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
