package composers

import (
	config "duna-pet-back/adapters/infrastructure/mySql/config"
	repository "duna-pet-back/adapters/infrastructure/mySql/repository/pets"
	interfaces "duna-pet-back/application/interfaces/pets"
	usecases "duna-pet-back/application/usecases/pets"
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
