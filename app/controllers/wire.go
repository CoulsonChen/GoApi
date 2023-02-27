//go:build wireinject

package controllers

import (
	"github.com/CoulsonChen/GoApi/app/services"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	services.ProviderSet,
	UsersControllerProvider,
)

func CreateUsersController() *UsersController {
	panic(wire.Build(ProviderSet))
}
