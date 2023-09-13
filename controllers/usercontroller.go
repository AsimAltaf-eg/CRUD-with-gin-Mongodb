package controllers

import (
	"main/models"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserServices services.UserServices
}

func New(userservice services.UserServices) UserController {
	return UserController{
		UserServices: userservice,
	}
}

func (uc *UserController) CreateUsers(ctx *gin.Context) {
	var Users *[]models.User

	if err := ctx.ShouldBindJSON(&Users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	if err := uc.UserServices.CreateUsers(Users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": "Users Created"})
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var NewUser models.User
	if err := ctx.ShouldBindJSON(&NewUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	if err := uc.UserServices.CreateUser(&NewUser); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Succes": "User Created"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {

	id := ctx.Param("name")
	if a, err := uc.UserServices.GetUser(&id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Error with Database Connection"})
		return
	} else {
		ctx.JSON(http.StatusOK, a)
	}

}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	allUsers, err := uc.UserServices.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Request is not Fulfilled"})
		return
	}
	ctx.JSON(http.StatusOK, allUsers)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	var newUser models.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Failed Updating the User"})
		return
	}
	err := uc.UserServices.UpdateUser(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Failed Updating the User"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Succes": "User Updated"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("name")

	if err := uc.UserServices.DeleteUser(&id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Error with Database Connection"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Success": "User Deleted"})
	}

}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroutes := rg.Group("/")
	userroutes.POST("/create", uc.CreateUser)
	userroutes.POST("/createAll", uc.CreateUsers)
	userroutes.GET("/", uc.GetUsers)
	userroutes.GET("/:name", uc.GetUser)
	userroutes.PATCH("/update", uc.UpdateUser)
	userroutes.DELETE("/delete/:name", uc.DeleteUser)
}
