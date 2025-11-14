package jobshandler

import (
	"github.com/alirezazahiri/gofetch-v2/internal/delivery/dto/jobsdto"
	"github.com/alirezazahiri/gofetch-v2/pkg/envelope"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Check(c *gin.Context) {
	var request jobsdto.CheckRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		envelope.BadRequest(c, "Invalid request format", err.Error())
		return
	}

	if len(request.Urls) == 0 {
		envelope.ValidationError(c, "Validation failed", map[string]string{
			"urls": "must not be empty",
		})
		return
	}

	if request.Concurrency <= 0 {
		envelope.ValidationError(c, "Validation failed", map[string]string{
			"concurrency": "must be greater than 0",
		})
		return
	}

	// TODO: Implement actual job creation logic here
	envelope.OK(c, map[string]string{
		"message": "Check job is running ✅",
		"status":  "accepted",
	})
}
