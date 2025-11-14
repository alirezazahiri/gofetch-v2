package jobsutils

import (
	"runtime"

	"github.com/alirezazahiri/gofetch-v2/internal/entity"
)

func MapJobStatusToString(status entity.JobStatus) string {
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

func MapJobResultStatusToString(status entity.JobResultStatus) string {
	switch status {
	case entity.JobResultStatusCompleted:
		return "completed"
	case entity.JobResultStatusFailed:
		return "failed"
	case entity.JobResultStatusTimeout:
		return "timeout"
	}
	return "unknown"
}

func NormalizeNumberOfWorkers(numberOfWorkers int) int {
	if numberOfWorkers < 1 {
		return 1
	}
	if numberOfWorkers > runtime.NumCPU() {
		return runtime.NumCPU()
	}
	return numberOfWorkers
}