package main

import (
	"github.com/rafabcanedo/api-golang/src/controller"
	"github.com/rafabcanedo/api-golang/src/model/repository"
	"github.com/rafabcanedo/api-golang/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}