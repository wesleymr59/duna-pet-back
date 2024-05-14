package controllers

import (
	composers "duna-pet-back/application/composers/healthyCheck"

	"github.com/gin-gonic/gin"
)

type HealthyController struct {
	composer composers.HealthyComposer
}

func NewHealthyController() *HealthyController {
	return &HealthyController{
		composer: *composers.NewHealthyComposer(),
	}
}

func (c *HealthyController) HandleHealty(ctx *gin.Context) {
	service := c.composer.ComposeHelloService()
	message := service.GetHelloMessage()

	ctx.JSON(200, gin.H{"message": message})
}

func (h *HealthyController) HandleIndex(c *gin.Context) {
	service := h.composer.ComposeHelloService()
	message := service.GetIndexMessage()
	c.JSON(200, gin.H{"message": message})
}
