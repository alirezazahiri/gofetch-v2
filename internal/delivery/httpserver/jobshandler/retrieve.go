package jobshandler

import (
	"github.com/alirezazahiri/gofetch-v2/internal/delivery/dto/jobsdto"
	"github.com/alirezazahiri/gofetch-v2/pkg/envelope"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RetrieveJob(c *gin.Context) {
	request := jobsdto.RetrieveRequest{
		ID: c.Param("id"),
	}

	response, err := h.svc.Retrieve(c.Request.Context(), &request)
	
	if err != nil {
		envelope.InternalServerError(c, "Failed to retrieve job", err.Error())
		return
	}

	envelope.OK(c, response)
}
