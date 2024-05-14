package interfaces

import "duna-pet-back/domain/entities"

type UserRepository interface {
	FindByEmailAndPassword(email, password string) (entities.User, error)
}
