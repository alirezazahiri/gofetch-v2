package entity

import "time"

type JobStatus uint8

const (
	JobStatusPending JobStatus = iota
	JobStatusRunning
	JobStatusCompleted
	JobStatusFailed
)

type Job struct {
	ID         string      `gorm:"primaryKey"`
	Status     JobStatus   `gorm:"not null;default:0"`
	DurationMs int64       `gorm:"not null"`
	CreatedAt  time.Time   `gorm:"not null"`
	UpdatedAt  time.Time   `gorm:"not null"`
	JobResults []JobResult `gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type JobResultStatus uint8

const (
	JobResultStatusCompleted JobResultStatus = iota
	JobResultStatusFailed
	JobResultStatusTimeout
)

type JobResult struct {
	ID        string          `gorm:"primaryKey"`
	JobID     string          `gorm:"not null;index"`
	Job       Job             `gorm:"foreignKey:JobID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Url       string          `gorm:"not null"`
	Status    JobResultStatus `gorm:"not null;default:0"`
	LatencyMs int64           `gorm:"not null"`
	CreatedAt time.Time       `gorm:"not null"`
	UpdatedAt time.Time       `gorm:"not null"`
}
