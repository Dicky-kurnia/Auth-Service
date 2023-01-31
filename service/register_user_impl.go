package service

import (
	"service-producer/entity"
	"service-producer/model"
	"service-producer/repository"
)

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) RegisterUserService {
	return &userServiceImpl{repo: repo}
}

func (s userServiceImpl) RegisterUserServ(request model.RegistrerUserRequest) (entity.UserEntity, error) {

	user := entity.UserEntity{}
	user.Name = request.Name
	user.Occupation = request.Occupation
	user.Email = request.Email

	newUser, err := s.repo.SaveUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
