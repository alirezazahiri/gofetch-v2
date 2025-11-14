package httpserver

import (
	"fmt"
	"log"

	"github.com/alirezazahiri/gofetch-v2/internal/delivery/httpserver/jobshandler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server   *gin.Engine
	cfg      *Config
	handlers *Handlers
}

type Handlers struct {
	JobsHandler *jobshandler.Handler
}

func NewServer(cfg *Config, handlers *Handlers) *Server {
	return &Server{
		server: gin.Default(),
		cfg:    cfg,
		handlers: &Handlers{
			JobsHandler: handlers.JobsHandler,
		},
	}
}

func (s *Server) Start() {
	router := s.server.Group("/api")

	router.GET("/health", Healthcheck)

	// jobs api routes 
	jobsRouter := router.Group("/jobs")
	s.handlers.JobsHandler.RegisterRoutes(jobsRouter)

	err := s.server.Run(fmt.Sprintf(":%d", s.cfg.Port))

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
