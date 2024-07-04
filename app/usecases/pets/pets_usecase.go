package usecases

import (
	"duna-pet-back/app/entities"
	interfaces "duna-pet-back/app/interfaces/pets"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PetsUsecase struct {
	petRepository interfaces.PetRepository
}

func NewPetService(repo interfaces.PetRepository) *PetsUsecase {
	return &PetsUsecase{petRepository: repo}
}

func (s *PetsUsecase) GetDogRacesByName(c *gin.Context) (entities.Dog, error) {
	var dog entities.Dog
	log.Println("dsadasdsa")
	races := c.Query("race")

	userFromDB, err := s.petRepository.FindDogRacesByName(races)
	log.Println(userFromDB)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return dog, errors.New("binding GET Races failed")
	}

	return userFromDB, nil
}
