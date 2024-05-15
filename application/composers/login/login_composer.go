package composers

import (
	config "duna-pet-back/adapters/infrastructure/mySql/config"
	repository "duna-pet-back/adapters/infrastructure/mySql/repository/login"
	interfaces "duna-pet-back/application/interfaces/login"
	usecases "duna-pet-back/application/usecases/login"
)

type LoginComposer struct{}

func NewLoginComposer() *LoginComposer {
	return &LoginComposer{}
}

func (c *LoginComposer) ComposeLoginService() *usecases.LoginUsecase {
	db := config.DB
	userRepo := repository.NewMySQLUserRepository(db)
	return usecases.NewLoginService(userRepo)
}

func (c *LoginComposer) GetUserRepository() interfaces.UserRepository {
	db := config.DB
	return repository.NewMySQLUserRepository(db)
}
