// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package controllers

import (
	"github.com/CoulsonChen/GoApi/internal/app/services"
	"github.com/CoulsonChen/GoApi/internal/pkg/db"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateUsersController() *UsersController {
	gormDB := db.DBProvider()
	usersService := services.UsersServiceProvider(gormDB)
	usersController := UsersControllerProvider(usersService)
	return usersController
}

// wire.go:

var ProviderSet = wire.NewSet(services.ProviderSet, UsersControllerProvider)
