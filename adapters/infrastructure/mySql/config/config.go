package database

import (
	"duna-pet-back/application/utils"
	"duna-pet-back/domain/entities"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnDataBase() {

	env_data := utils.GetConfig()
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", env_data.Username, env_data.Password, env_data.Host, env_data.Port, env_data.Name)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com banco de dados")
	}
	conn.AutoMigrate(&entities.User{}, &entities.Dog{})

	DB = conn
}
