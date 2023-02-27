//go:build wireinject

package services

import (
	"github.com/CoulsonChen/GoApi/pkg/db"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.DBProvider,
	UsersServiceProvider,
	wire.Bind(new(IUsersService), new(*UsersService)),
)

func CreateUsersService() (*UsersService, error) {
	panic(wire.Build(ProviderSet))
}
