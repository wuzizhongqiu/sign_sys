// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/conf"
	"wuzigoweb/internal/data"
	"wuzigoweb/internal/server"
	"wuzigoweb/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, micro *conf.Micro, logger log.Logger) (*kratos.App, func(), error) {
	db, err := data.NewDB(confData)
	if err != nil {
		return nil, nil, err
	}
	client, err := data.NewRedis(confData)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(db, logger, client)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	userService := service.NewUserService(userUsecase)
	classRepo := data.NewClassRepo(dataData, logger)
	classUsecase := biz.NewClassUsecase(classRepo, logger)
	classService := service.NewClassService(classUsecase)
	httpServer := server.NewHTTPServer(confServer, userService, classService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
