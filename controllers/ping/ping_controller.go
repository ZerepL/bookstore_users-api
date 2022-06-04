package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Return pong
// @Description  return a pong in order to test health
// @Tags         ping
// @Success      200  "pong"
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
