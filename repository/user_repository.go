package repository

import (
	"go-clean-architecture/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(u *entity.User, email string) error
	CreateUser(u *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(u *entity.User, email string) error {
	if err := ur.db.Where("email=?", email).First(u).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(u *entity.User) error {
	if err := ur.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}
