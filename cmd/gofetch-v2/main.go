package main

import (
	"github.com/alirezazahiri/gofetch-v2/internal/config"
	"github.com/alirezazahiri/gofetch-v2/internal/delivery/httpserver"
	"github.com/alirezazahiri/gofetch-v2/internal/delivery/httpserver/jobshandler"
)


func main() {
	cfg := config.Load("config.yml")

	jobsHandler := jobshandler.New()
	server := httpserver.NewServer(&cfg.HttpServer, &httpserver.Handlers{
		JobsHandler: jobsHandler,
	})
	
	server.Start()
}