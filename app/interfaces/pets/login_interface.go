package interfaces

import "duna-pet-back/app/entities"

type PetRepository interface {
	FindDogRacesByName(race string) (entities.Dog, error)
}
