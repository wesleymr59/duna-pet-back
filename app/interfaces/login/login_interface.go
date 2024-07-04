package interfaces

import "duna-pet-back/app/entities"

type UserRepository interface {
	FindByEmailAndPassword(email, password string) (entities.User, error)
	CreateUser(user *entities.User) error
}
