package worker

import (
	"context"
	"log"
	"sync"
)

// worker processes jobs from the jobs channel and sends results to the results channel
func worker[T any, R any](ctx context.Context, wg *sync.WaitGroup, jobs <-chan T, results chan<- R, processFn func(T) R) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			result := processFn(job)
			results <- result
		}
	}
}

// Run creates a worker pool that processes jobs concurrently
// T is the type of input jobs, R is the type of results
// It returns a channel that will receive all results
func Run[T any, R any](ctx context.Context, jobs []T, workers int, processFn func(T) R) <-chan R {
	in := make(chan T)
	out := make(chan R)
	var wg sync.WaitGroup

	// Start worker pool
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, in, out, processFn)
	}

	// Feed jobs and close channels when done
	go func() {
		for _, job := range jobs {
			select {
			case <-ctx.Done():
				close(in)
				return
			case in <- job:
				log.Printf("JOB_SENT: job=%v", job)
			}
		}
		close(in)
		wg.Wait()
		close(out)
	}()

	return out
}
