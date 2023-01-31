package service

import (
	"service-producer/entity"
	"service-producer/model"
)

type RegisterUserService interface {
	RegisterUserServ(request model.RegistrerUserRequest) (entity.UserEntity, error)
}
