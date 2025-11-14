package jobsservice

import (
	"context"
	"log"
	"time"

	"github.com/alirezazahiri/gofetch-v2/internal/delivery/dto/jobsdto"
	"github.com/alirezazahiri/gofetch-v2/internal/entity"
	"github.com/alirezazahiri/gofetch-v2/internal/repository/postgresql/jobsrepo"
	"github.com/alirezazahiri/gofetch-v2/pkg/pinger"
	"github.com/alirezazahiri/gofetch-v2/pkg/uuid"
	"github.com/alirezazahiri/gofetch-v2/pkg/worker"
)

type Service struct {
	repo *jobsrepo.Repository
}

func New(repo *jobsrepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Check(ctx context.Context, request *jobsdto.CheckRequest) (*jobsdto.CheckResponse, error) {
	job := &entity.Job{
		ID:        uuid.New(),
		Status:    entity.JobStatusPending,
		DurationMs: 0,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.repo.CreateJob(job); err != nil {
		log.Println("CREATE_JOB_ERROR:", job.ID, err)
		return nil, err
	}

	// run the workers with the number of `request.Concurrency` goroutines
	// each worker will have a timeout of `request.Timeout`
	for i := 0; i < request.Concurrency; i++ {
		for _, url := range request.Urls {
			go func(jobID, url string) {
				worker.StartWorker(ctx, func() error {
					start := time.Now()
					err := pinger.Ping(url)
					duration := time.Since(start)
					go func(jobID, url string) {
						jobResult := &entity.JobResult{
							ID:        uuid.New(),
							JobID:     jobID,
							Url:       url,
							Status:    entity.JobResultStatusCompleted,
							LatencyMs: duration.Milliseconds(),
							CreatedAt: time.Now().UTC(),
							UpdatedAt: time.Now().UTC(),
						}
						if err != nil {
							jobResult.Status = entity.JobResultStatusFailed
						}
						createError := s.repo.CreateJobResult(jobResult)
						if createError != nil {
							log.Println("CREATE_JOB_RESULT_ERROR:", jobID, url, createError)
						}
					}(jobID, url)
					return nil
				}, time.Duration(request.Timeout)*time.Second)
			}(job.ID, url)
		}
	}

	return &jobsdto.CheckResponse{
		JobId:  job.ID,
		Status: mapJobStatusNumberToString(job.Status),
	}, nil
}

func mapJobStatusNumberToString(status entity.JobStatus) string {
	switch status {
	case entity.JobStatusPending:
		return "pending"
	case entity.JobStatusRunning:
		return "running"
	case entity.JobStatusCompleted:
		return "completed"
	case entity.JobStatusFailed:
		return "failed"
	}
	return "unknown"
}
