package repository

import (
	"duna-pet-back/domain/entities"

	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(user entities.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySQLUserRepository) FindByEmailAndPassword(email, password string) (entities.User, error) {
	var user entities.User
	if result := r.db.Where("email = ? AND passwd = ? AND active = ?", email, password, true).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
