package jobsdto

import "time"

type RetrieveRequest struct {
	ID string `json:"id"`
}

type JobResultItem struct {
	URL       string `json:"url"`
	LatencyMs int64  `json:"latency_ms"`
	Status    string `json:"status"`
}

type RetrieveResponse struct {
	JobID      string          `json:"job_id"`
	Results    []JobResultItem `json:"results"`
	Status     string          `json:"status"`
	DurationMs int64           `json:"duration_ms"`
	CreatedAt  time.Time       `json:"created_at"`
}
