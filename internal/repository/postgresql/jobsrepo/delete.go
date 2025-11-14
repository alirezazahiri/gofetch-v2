package jobsrepo

import "github.com/alirezazahiri/gofetch-v2/internal/entity"

func (r *Repository) DeleteJob(id string) error {
	job := &entity.Job{
		ID: id,
	}
	return r.db.Where("id = ?", id).Delete(job).Error
}

func (r *Repository) DeleteJobResult(id string) error {
	jobResult := &entity.JobResult{
		ID: id,
	}
	return r.db.Where("id = ?", id).Delete(jobResult).Error
}
