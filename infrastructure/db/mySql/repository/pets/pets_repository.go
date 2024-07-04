package repository

import (
	"duna-pet-back/app/entities"

	"gorm.io/gorm"
)

type MySQLPetsRepository struct {
	db *gorm.DB
}

func NewMySQLPetsRepository(db *gorm.DB) *MySQLPetsRepository {
	return &MySQLPetsRepository{db: db}
}

func (repo *MySQLPetsRepository) FindDogRacesByName(race string) (entities.Dog, error) {
	var dog entities.Dog
	if err := repo.db.Where("Race = ?", race).First(&dog).Error; err != nil {
		return entities.Dog{}, err
	}
	return dog, nil
}
