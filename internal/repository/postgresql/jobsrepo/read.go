package jobsrepo

import "github.com/alirezazahiri/gofetch-v2/internal/entity"

func (r *Repository) GetJob(id string) (*entity.Job, error) {
	job := &entity.Job{}
	return job, r.db.Where("id = ?", id).First(job).Error
}

func (r *Repository) GetJobWithResults(jobID string) (*entity.Job, error) {
	job := &entity.Job{}
	return job, r.db.Preload("JobResults").Where("id = ?", jobID).First(job).Error
}

func (r *Repository) GetJobResult(id string) (*entity.JobResult, error) {
	jobResult := &entity.JobResult{}
	return jobResult, r.db.Where("id = ?", id).First(jobResult).Error
}
