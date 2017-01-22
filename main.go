package main

import (
	"context"
	"fmt"
	"time"
)

func ShowValue(ctx context.Context, key string) {
	fmt.Println(time.Now(), ctx.Value(key))
}

func Perform(ctx context.Context) error {
	for {
		ShowValue(ctx, "custom-key")

		select {
		case <-ctx.Done():
			err := ctx.Err()

			// Show error value before returning it.
			fmt.Println(err)

			return err
		case <-time.After(time.Second):
			continue
		}
	}
	return nil
}

func DoCancel(cancel context.CancelFunc, duration time.Duration) {
	time.Sleep(duration)
	cancel()
}

func main() {
	// Create Context with key "custom-key" and value "custom-value"
	ctx := context.WithValue(context.Background(), "custom-key", "custom value")

	// Set 5-second timeout.
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	go Perform(ctx)

	// Trigger cancel after 3 seconds.
	go DoCancel(cancel, time.Second*3)

	time.Sleep(time.Second * 10)
}
