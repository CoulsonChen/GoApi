package services

import (
	"github.com/CoulsonChen/GoApi/internal/pkg/models"
	"gorm.io/gorm"
)

type IUsersService interface {
	GetAllUsers() (*[]models.User, error)
	GetUserByFullName(fullName string) (user *models.User, err error)
}

type UsersService struct {
	db *gorm.DB
}

func UsersServiceProvider(db *gorm.DB) *UsersService {
	return &UsersService{
		db: db,
	}
}

func (service *UsersService) GetAllUsers() (users *[]models.User, err error) {
	service.db.Find(&users)
	err = nil
	return
}

func (service *UsersService) GetUserByFullName(fullName string) (user *models.User, err error) {
	service.db.Where("fullname = ?", fullName).First(&user)
	err = nil
	return
}
