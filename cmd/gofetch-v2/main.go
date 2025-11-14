package main

import (
	"log"

	"github.com/alirezazahiri/gofetch-v2/internal/config"
	"github.com/alirezazahiri/gofetch-v2/internal/delivery/httpserver"
	"github.com/alirezazahiri/gofetch-v2/internal/delivery/httpserver/jobshandler"
	"github.com/alirezazahiri/gofetch-v2/internal/jobsservice"
	"github.com/alirezazahiri/gofetch-v2/internal/repository/postgresql"
	"github.com/alirezazahiri/gofetch-v2/internal/repository/postgresql/jobsrepo"
)


func main() {
	cfg := config.Load("config.yml")

	postgresRepo, err := postgresql.New(&cfg.Repository.Postgres)
	if err != nil {
		log.Fatalf("failed to create postgres repository: %v", err)
	}
	defer postgresRepo.Close()

	jobsRepo := jobsrepo.New(postgresRepo.DB())
	jobsService := jobsservice.New(jobsRepo)
	jobsHandler := jobshandler.New(jobsService)

	server := httpserver.NewServer(&cfg.HttpServer, &httpserver.Handlers{
		JobsHandler: jobsHandler,
	})

	server.Start()
}
