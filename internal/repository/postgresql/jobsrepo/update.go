package jobsrepo

import "github.com/alirezazahiri/gofetch-v2/internal/entity"

func (r *Repository) UpdateJob(job *entity.Job) error {
	return r.db.Save(job).Error
}
