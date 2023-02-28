package services

import (
	"errors"

	"github.com/CoulsonChen/GoApi/pkg/models"
	"gorm.io/gorm"
)

type IUsersService interface {
	GetAllUsers() (*[]models.User, error)
	GetUserByFullName(fullName string) (user *models.User, err error)
	CreateUser(user models.User) (useracct *string, err error)
	Login(login models.Login) (isSuccess *bool, err error)
	DeleteUserByAccount(acct string) (user *models.User, err error)
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

func (service *UsersService) CreateUser(user models.User) (useracct *string, err error) {
	service.db.Create(&user)
	useracct = &user.Acct
	err = nil
	return
}

func (service *UsersService) Login(login models.Login) (isSuccess *bool, err error) {
	result := false
	err = nil

	user := models.User{}
	if err := service.db.Table("users").Where("acct=?", login.Acct).First(&user).Error; err != nil {
		return &result, err
	}

	if user.Acct == "" {
		return &result, errors.New("account not exist")
	}

	if user.Pwd != login.Pwd {
		return &result, errors.New("password incorrect")
	}

	result = true
	return &result, nil
}

func (service *UsersService) DeleteUserByAccount(acct string) (user *models.User, err error) {
	service.db.Table("users").Where("acct = ?", acct).Delete(&user)
	err = nil
	return
}
