package composers

import (
	interfaces "duna-pet-back/app/interfaces/login"
	usecases "duna-pet-back/app/usecases/login"
	config "duna-pet-back/infrastructure/db/mySql/config"
	repository "duna-pet-back/infrastructure/db/mySql/repository/login"
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
