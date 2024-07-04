package main

import (
	loginComposers "duna-pet-back/app/composers/login"
	petComposers "duna-pet-back/app/composers/pets"
	"duna-pet-back/app/routers"
	config "duna-pet-back/infrastructure/db/mySql/config"
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
