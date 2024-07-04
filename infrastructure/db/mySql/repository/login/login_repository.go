package repository

import (
	"duna-pet-back/app/entities"

	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (repo *MySQLUserRepository) FindByEmailAndPassword(email, password string) (entities.User, error) {
	var user entities.User
	if err := repo.db.Where("email = ? AND passwd = ?", email, password).First(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (repo *MySQLUserRepository) CreateUser(user *entities.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
