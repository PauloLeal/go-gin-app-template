package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (ctrl *HealthController) Get(c *gin.Context) {
	d := map[string]interface{}{"status": "ok"}
	c.JSON(200, d)
}
