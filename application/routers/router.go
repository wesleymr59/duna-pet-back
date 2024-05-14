package routers

import (
	HealthyController "duna-pet-back/application/controllers/healthyCheck"
	loginController "duna-pet-back/application/controllers/login"
	"duna-pet-back/application/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	helloController := HealthyController.NewHealthyController()
	loginController := loginController.NewLoginHandler()

	router.GET("/hello", helloController.HandleHealty)
	router.GET("/", helloController.HandleIndex)

	router.GET("/login", loginController.Login)
	router.POST("/register", loginController.RegisterHandler)
	router.GET("/refresh-token", utils.RefreshToken)

	return router
}
