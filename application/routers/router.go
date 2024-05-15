package routers

import (
	healthyControllers "duna-pet-back/application/controllers/healthyCheck"
	loginControllers "duna-pet-back/application/controllers/login"
	petControllers "duna-pet-back/application/controllers/pets"
	loginInterfaces "duna-pet-back/application/interfaces/login"
	petInterfaces "duna-pet-back/application/interfaces/pets"
	loginUsecases "duna-pet-back/application/usecases/login"
	petUsecases "duna-pet-back/application/usecases/pets"
	"duna-pet-back/application/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userRepo loginInterfaces.UserRepository, petRepo petInterfaces.PetRepository) *gin.Engine {
	router := gin.Default()

	helloController := healthyControllers.NewHealthyController()
	loginService := loginUsecases.NewLoginService(userRepo)
	loginController := loginControllers.NewLoginHandler(loginService)

	petService := petUsecases.NewPetService(petRepo)
	petController := petControllers.NewPetsHandler(petService)

	router.GET("/hello", helloController.HandleHealty)
	router.GET("/", helloController.HandleIndex)

	router.POST("/login", loginController.Login)
	router.POST("/register", loginController.RegisterHandler)
	router.GET("/refresh-token", utils.RefreshToken)

	router.GET("/races", petController.GetDogRacesByName)
	return router
}
