package worker

import (
	"context"
	"log"
	"time"
)

func worker(ctx context.Context, cb func() error, timeout time.Duration) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := cb()
			if err != nil {
				return err
			}
			return nil
		}
	}
}

func StartWorker(ctx context.Context, cb func() error, timeout time.Duration) {
	go func() {
		err := worker(ctx, cb, timeout)
		if err != nil {
			log.Printf("WORKER_ERROR: %v", err)
		}
	}()
}
