package jobsrepo

import "github.com/alirezazahiri/gofetch-v2/internal/entity"

func (r *Repository) CreateJob(job *entity.Job) error {
	return r.db.Create(job).Error
}

func (r *Repository) CreateJobResult(jobResult *entity.JobResult) error {
	return r.db.Create(jobResult).Error
}
