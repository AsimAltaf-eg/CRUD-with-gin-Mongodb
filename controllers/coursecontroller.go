package controllers

import (
	"main/models"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	CourseService services.CourseService
}

func NewCourseController(courseservice services.CourseService) CourseController {
	return CourseController{
		CourseService: courseservice,
	}
}

func (uc *CourseController) CreateCourse(ctx *gin.Context) {
	var newCourse models.Course
	if err := ctx.ShouldBindJSON(&newCourse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Type"})
		return
	}

	if err := uc.CourseService.CreateCourse(&newCourse); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Bad Gateway"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": "Course Created"})

}

func (uc *CourseController) GetCourses(ctx *gin.Context) {

	if allCourses, err := uc.CourseService.GetCourses(); err != nil {

		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Bad Gateway"})
		return
	} else {
		ctx.JSON(http.StatusOK, allCourses)
	}
}

func (uc *CourseController) UpdateCourse(ctx *gin.Context) {

	id := ctx.Param("id")
	var a models.Course
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Bad Object"})
		return
	}

	if err := uc.CourseService.UpdateCourse(&id, &a); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Bad Gateway"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": "Course Updated"})
}

func (uc *CourseController) DeleteCourse(ctx *gin.Context) {

	id := ctx.Param("code")
	if err := uc.CourseService.DeleteCourse(&id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Bad Gateway"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": "Course Deleted"})
}
func (uc *CourseController) AddCourses(ctx *gin.Context) {

	var NewCourse *[]models.Course
	if err := ctx.ShouldBindJSON(&NewCourse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Type"})
		return
	}

	if err := uc.CourseService.AddCourses(NewCourse); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Bad Gateway"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": "Courses Created"})
}
func (uc *CourseController) GetCourse(ctx *gin.Context) {

	id := ctx.Param("id")
	if course, err := uc.CourseService.GetCourse(&id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Bad Gateway"})
		return
	} else {

		ctx.JSON(http.StatusOK, course)
	}

}

func (uc *CourseController) RegisterCourseRoutes(rg *gin.RouterGroup) {
	userroutes := rg.Group("/")
	userroutes.POST("/create", uc.CreateCourse)
	userroutes.POST("/createAll", uc.AddCourses)
	userroutes.GET("/", uc.GetCourses)
	userroutes.GET("/:id", uc.GetCourse)
	userroutes.PATCH("/update/:id", uc.UpdateCourse)
	userroutes.DELETE("/delete/:code", uc.DeleteCourse)
}
