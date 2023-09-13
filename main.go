package main

import (
	"context"
	"log"
	"main/controllers"
	"main/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx            context.Context
	server         *gin.Engine
	userService    services.UserServices
	userController controllers.UserController
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal("Applicaton Failed to Connect to the Database")
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Applicaton Failed to Connect to the Database")
	}

	usercollection = mongoclient.Database("local").Collection("users")
	userService = services.NewUserService(usercollection, ctx)
	userController = controllers.New(userService)

	server = gin.Default()
}

func main() {

	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/home")
	userController.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":8080"))
}
