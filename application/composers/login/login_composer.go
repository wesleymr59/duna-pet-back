package composers

import (
	usecases "duna-pet-back/application/usecases/login"
)

type LoginComposer struct{}

func NewLoginComposer() *LoginComposer {
	return &LoginComposer{}
}

func (c *LoginComposer) ComposeLoginService() *usecases.LoginUsecase {
	return usecases.NewLoginService()
}
