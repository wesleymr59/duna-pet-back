package interfaces

import (
	composers "duna-pet-back/application/composers/login"
	"duna-pet-back/domain/entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	composer composers.LoginComposer
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{
		composer: *composers.NewLoginComposer(),
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var login entities.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.LoginUser(login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *LoginHandler) RegisterHandler(c *gin.Context) {
	service := h.composer.ComposeLoginService()
	user, _ := service.RegisterUser(c)
	fmt.Println(user)
	c.JSON(http.StatusOK, user)
}
