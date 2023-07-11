// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"mall-shop-be/internal/biz"
	"mall-shop-be/internal/conf"
	"mall-shop-be/internal/data"
	"mall-shop-be/internal/errors"
	"mall-shop-be/internal/server"
	"mall-shop-be/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedisConn(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	errorController := errors.NewErrorController()
	userRepo := data.NewUserRepo(dataData, logger, errorController)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase)
	grpcServer := server.NewGRPCServer(confServer, userService, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
