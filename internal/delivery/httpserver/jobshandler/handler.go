package jobshandler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/check", h.Check)
}