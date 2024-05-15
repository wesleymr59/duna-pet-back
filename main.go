package main

import (
	config "duna-pet-back/adapters/infrastructure/mySql/config"
	loginComposers "duna-pet-back/application/composers/login"
	petComposers "duna-pet-back/application/composers/pets"
	"duna-pet-back/application/routers"
)

func main() {
	// Conectar ao banco de dados
	config.ConnDataBase()

	// Compositor para criar o serviço de login e repositório
	composerLogin := loginComposers.NewLoginComposer()
	composerPet := petComposers.NewPetsComposer()
	// Configurar o roteador com o repositório de usuários injetado
	router := routers.SetupRouter(composerLogin.GetUserRepository(), composerPet.GetPetRepository())

	// Iniciar o servidor
	router.Run(":3001")
}
