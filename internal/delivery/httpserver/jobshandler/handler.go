package jobshandler

import (
	"github.com/alirezazahiri/gofetch-v2/internal/jobsservice"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *jobsservice.Service
}

func New(svc *jobsservice.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/check", h.Check)
	router.GET("/:id", h.RetrieveJob)
}