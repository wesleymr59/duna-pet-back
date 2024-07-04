package composers

import (
	interfaces "duna-pet-back/app/interfaces/pets"
	usecases "duna-pet-back/app/usecases/pets"
	config "duna-pet-back/infrastructure/db/mySql/config"
	repository "duna-pet-back/infrastructure/db/mySql/repository/pets"
)

type PetsComposer struct{}

func NewPetsComposer() *PetsComposer {
	return &PetsComposer{}
}

func (c *PetsComposer) ComposePetsService() *usecases.PetsUsecase {
	db := config.DB
	userRepo := repository.NewMySQLPetsRepository(db)
	return usecases.NewPetService(userRepo)
}
func (c *PetsComposer) GetPetRepository() interfaces.PetRepository {
	db := config.DB
	return repository.NewMySQLPetsRepository(db)
}
