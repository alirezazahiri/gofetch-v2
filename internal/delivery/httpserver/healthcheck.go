package httpserver

import (
	"github.com/alirezazahiri/gofetch-v2/pkg/envelope"
	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	envelope.OK(c, map[string]string{
		"status":  "healthy",
		"message": "Server is Up and Running ✅",
	})
}
