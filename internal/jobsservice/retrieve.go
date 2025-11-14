package jobsservice

import (
	"context"

	"github.com/alirezazahiri/gofetch-v2/internal/delivery/dto/jobsdto"
	"github.com/alirezazahiri/gofetch-v2/pkg/jobsutils"
)

func (s *Service) Retrieve(ctx context.Context, request *jobsdto.RetrieveRequest) (*jobsdto.RetrieveResponse, error) {
	job, err := s.repo.GetJobWithResults(request.ID)
	if err != nil {
		return nil, err
	}

	results := make([]jobsdto.JobResultItem, len(job.JobResults))
	for i, result := range job.JobResults {
		results[i] = jobsdto.JobResultItem{
			URL:       result.Url,
			LatencyMs: result.LatencyMs,
			Status:    jobsutils.MapJobResultStatusToString(result.Status),
		}
	}

	return &jobsdto.RetrieveResponse{
		JobID:      job.ID,
		DurationMs: job.DurationMs,
		CreatedAt:  job.CreatedAt,
		Results:    results,
		Status:     jobsutils.MapJobStatusToString(job.Status),
	}, nil
}
