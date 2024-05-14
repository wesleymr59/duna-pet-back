package main

import (
	database "duna-pet-back/adapters/infrastructure/mySql/config"
	"duna-pet-back/application/routers"
	"log"
)

func main() {
	database.ConnDataBase()
	log.Println("Service Starting....")
	routers.SetupRouter().Run(":8080")
}
