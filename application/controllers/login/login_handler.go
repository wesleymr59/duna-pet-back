package controllers

import (
	usecases "duna-pet-back/application/usecases/login"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginService *usecases.LoginUsecase
}

func NewLoginHandler(loginService *usecases.LoginUsecase) *LoginHandler {
	return &LoginHandler{
		loginService: loginService,
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	token, err := h.loginService.LoginUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *LoginHandler) RegisterHandler(c *gin.Context) {
	h.loginService.RegisterUser(c)
	c.JSON(http.StatusOK, nil)
}
