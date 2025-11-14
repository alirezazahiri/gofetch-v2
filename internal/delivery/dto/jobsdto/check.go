package jobsdto

import "time"

type CheckRequest struct {
	Urls        []string `json:"urls"`
	Concurrency int      `json:"concurrency"`
	TimeoutMs   int      `json:"timeout_ms"`
}

type CheckResponse struct {
	JobId     string    `json:"job_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
