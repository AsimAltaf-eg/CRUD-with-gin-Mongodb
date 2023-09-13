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
	ctx              context.Context
	server           *gin.Engine
	userService      services.UserServices
	userController   controllers.UserController
	usercollection   *mongo.Collection
	mongoclient      *mongo.Client
	courseService    services.CourseService
	courseController controllers.CourseController
	coursecollection *mongo.Collection
	err              error
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

	server = gin.Default()

	usercollection = mongoclient.Database("local").Collection("users")
	userService = services.NewUserService(usercollection, ctx)
	userController = controllers.NewUserController(userService)
	coursecollection = mongoclient.Database("local").Collection("courses")
	courseService = services.NewCourseService(coursecollection, ctx)
	courseController = controllers.NewCourseController(courseService)

}

func main() {

	basepath := server.Group("/user")
	userController.RegisterUserRoutes(basepath)

	defer mongoclient.Disconnect(ctx)
	basepathcourse := server.Group("/course")
	courseController.RegisterCourseRoutes(basepathcourse)

	log.Fatal(server.Run(":8080"))
}
