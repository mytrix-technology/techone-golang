package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration1 = 1 * time.Millisecond

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration1)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
