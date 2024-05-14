package interfaces

import (
	composers "duna-pet-back/application/composers/healthyCheck"

	"github.com/gin-gonic/gin"
)

type HealthyHandler struct {
	composer composers.HealthyComposer
}

func NewHealthyHandler() *HealthyHandler {
	return &HealthyHandler{
		composer: *composers.NewHealthyComposer(),
	}
}

func (h *HealthyHandler) HandleHello(c *gin.Context) {
	service := h.composer.ComposeHelloService()
	message := service.GetHelloMessage()
	c.JSON(200, gin.H{"message": message})
}

func (h *HealthyHandler) HandleIndex(c *gin.Context) {
	service := h.composer.ComposeHelloService()
	message := service.GetIndexMessage()
	c.JSON(200, gin.H{"message": message})
}
