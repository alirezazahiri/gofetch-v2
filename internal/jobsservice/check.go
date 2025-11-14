package jobsservice

import (
	"context"
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/alirezazahiri/gofetch-v2/internal/delivery/dto/jobsdto"
	"github.com/alirezazahiri/gofetch-v2/internal/entity"
	"github.com/alirezazahiri/gofetch-v2/pkg/jobsutils"
	"github.com/alirezazahiri/gofetch-v2/pkg/pinger"
	"github.com/alirezazahiri/gofetch-v2/pkg/uuid"
	"github.com/alirezazahiri/gofetch-v2/pkg/worker"
)

type pingResult struct {
	url       string
	latencyMs int64
	err       error
}

const (
	TimeoutError = "timeout"
)

func (s *Service) Check(ctx context.Context, request *jobsdto.CheckRequest) (*jobsdto.CheckResponse, error) {
	job := &entity.Job{
		ID:         uuid.New(),
		Status:     entity.JobStatusPending,
		DurationMs: 0,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	if err := s.repo.CreateJob(job); err != nil {
		log.Println("CREATE_JOB_ERROR:", job.ID, err)
		return nil, err
	}

	numberOfWorkers := jobsutils.NormalizeNumberOfWorkers(request.Concurrency)

	go func(jobID string) {
		asyncCtx := context.Background()
		start := time.Now()
		var isRunning atomic.Bool

		results := worker.Run(asyncCtx, request.Urls, numberOfWorkers, func(url string) pingResult {
			if !isRunning.Swap(true) {
				log.Printf("JOB_RUNNING: job=%s", jobID)
				job.Status = entity.JobStatusRunning
				job.UpdatedAt = time.Now().UTC()
				if err := s.repo.UpdateJob(job); err != nil {
					log.Println("UPDATE_JOB_TO_RUNNING_ERROR:", jobID, err)
				}
			}

			pingStart := time.Now()

			pingCtx, cancel := context.WithTimeout(asyncCtx, time.Duration(request.TimeoutMs)*time.Millisecond)
			defer cancel()

			errChan := make(chan error, 1)
			go func() {
				errChan <- pinger.Ping(url)
			}()

			var pingErr error
			select {
			case <-pingCtx.Done():
				pingErr = pingCtx.Err()
				if pingErr == context.DeadlineExceeded {
					pingErr = errors.New(TimeoutError)
				}
			case pingErr = <-errChan:
				log.Printf("PING_ERROR: job=%s url=%s error=%v", jobID, url, pingErr)
			}

			latency := time.Since(pingStart)

			return pingResult{
				url:       url,
				latencyMs: latency.Milliseconds(),
				err:       pingErr,
			}
		})

		countErrors := 0
		wg := sync.WaitGroup{}
		for result := range results {
			wg.Add(1)
			go func(result pingResult) {
				defer wg.Done()
				jobResult := &entity.JobResult{
					ID:        uuid.New(),
					JobID:     jobID,
					Url:       result.url,
					Status:    entity.JobResultStatusCompleted,
					LatencyMs: result.latencyMs,
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				}

				if result.err != nil {
					if result.err.Error() == TimeoutError {
						jobResult.Status = entity.JobResultStatusTimeout
					} else {
						jobResult.Status = entity.JobResultStatusFailed
						jobResult.LatencyMs = 0
					}
					countErrors++
					log.Printf("\n\nPING_ERROR: job=%s url=%s error=%v", jobID, result.url, result.err)
				}

				if err := s.repo.CreateJobResult(jobResult); err != nil {
					log.Println("\n\nCREATE_JOB_RESULT_ERROR:", jobID, result.url, err)
				}
			}(result)
		}

		wg.Wait()

		duration := time.Since(start)
		job.DurationMs = duration.Milliseconds()
		job.UpdatedAt = time.Now().UTC()

		if countErrors == len(request.Urls) {
			job.Status = entity.JobStatusFailed
			log.Printf("\n\nJOB_FAILED: job=%s countErrors=%d", jobID, countErrors)
		} else {
			job.Status = entity.JobStatusCompleted
			log.Printf("\n\nJOB_COMPLETED: job=%s countErrors=%d", jobID, countErrors)
		}

		log.Printf("\n\nUPDATE_JOB_FINAL_STATUS: job=%s status=%s countErrors=%d", jobID, jobsutils.MapJobStatusToString(job.Status), countErrors)
		if err := s.repo.UpdateJob(job); err != nil {
			log.Println("UPDATE_JOB_FINAL_STATUS_ERROR:", jobID, err)
		}
	}(job.ID)

	return &jobsdto.CheckResponse{
		JobId:  job.ID,
		Status: jobsutils.MapJobStatusToString(job.Status),
	}, nil
}
