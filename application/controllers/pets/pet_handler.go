package controllers

import (
	usecases "duna-pet-back/application/usecases/pets"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PetHandler struct {
	petService *usecases.PetsUsecase
}

func NewPetsHandler(petService *usecases.PetsUsecase) *PetHandler {
	return &PetHandler{
		petService: petService,
	}
}

func (h *PetHandler) GetDogRacesByName(c *gin.Context) {
	races, err := h.petService.GetDogRacesByName(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": races})
}
