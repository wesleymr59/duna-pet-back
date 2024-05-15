package interfaces

import "duna-pet-back/domain/entities"

type PetRepository interface {
	FindDogRacesByName(race string) (entities.Dog, error)
}
