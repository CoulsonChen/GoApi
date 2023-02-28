//go:build wireinject

package main

import (
	"github.com/CoulsonChen/GoApi/api"
	"github.com/CoulsonChen/GoApi/api/controllers"
	"github.com/CoulsonChen/GoApi/api/middlewares"
	"github.com/CoulsonChen/GoApi/pkg/config"
	"github.com/CoulsonChen/GoApi/pkg/db"
	"github.com/CoulsonChen/GoApi/pkg/services"
	"github.com/google/wire"
)

var (
	controllerProviderSet = wire.NewSet(
		controllers.UsersControllerProvider,
	)
	middlewareProviderSet = wire.NewSet(
		middlewares.JwtMiddlewareProvider,
	)
	serviceProviderSet = wire.NewSet(
		services.UsersServiceProvider,
		wire.Bind(new(services.IUsersService), new(*services.UsersService)),
		services.JwtServiceProvider,
		wire.Bind(new(services.IJwtService), new(*services.JwtService)),
	)
	configProviderSet = wire.NewSet(
		config.JwtConfigProvider,
		config.DbConfigProvider,
	)
)

var providerSet = wire.NewSet(
	db.DBProvider,
	route.RouteProvider,
	controllerProviderSet,
	middlewareProviderSet,
	serviceProviderSet,
	configProviderSet,
)

func InitHost() *route.Route {
	panic(wire.Build(providerSet))
}
