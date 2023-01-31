package repository

import (
	"service-producer/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo userRepository) SaveUser(user entity.UserEntity) (entity.UserEntity, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}
