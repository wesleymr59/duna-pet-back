package routers

import (
	healthyControllers "duna-pet-back/app/controllers/healthyCheck"
	loginControllers "duna-pet-back/app/controllers/login"
	petControllers "duna-pet-back/app/controllers/pets"
	loginInterfaces "duna-pet-back/app/interfaces/login"
	petInterfaces "duna-pet-back/app/interfaces/pets"
	loginUsecases "duna-pet-back/app/usecases/login"
	petUsecases "duna-pet-back/app/usecases/pets"
	"duna-pet-back/app/utils"

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
