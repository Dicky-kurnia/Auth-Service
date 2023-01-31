package repository

import "service-producer/entity"

type UserRepository interface {
	SaveUser(user entity.UserEntity) (entity.UserEntity, error)
}
