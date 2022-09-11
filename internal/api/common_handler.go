package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonHandler struct{}

func NewCommonHandler(g *gin.Engine) {
	handler := CommonHandler{}
	g.GET("/healthz", handler.HealthCheck)
}

func (h *CommonHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, Ok(nil, "Ok"))
}
