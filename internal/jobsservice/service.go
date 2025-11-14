package jobsservice

import (
	"github.com/alirezazahiri/gofetch-v2/internal/repository/postgresql/jobsrepo"
)

type Service struct {
	repo *jobsrepo.Repository
}

func New(repo *jobsrepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
