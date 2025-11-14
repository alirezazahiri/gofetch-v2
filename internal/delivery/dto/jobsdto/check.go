package jobsdto

import "time"

type CheckRequest struct {
	Urls        []string `json:"urls"`
	Concurrency int      `json:"concurrency"`
	Timeout     int      `json:"timeout"`
}

type CheckResponse struct {
	JobId     string    `json:"job_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
