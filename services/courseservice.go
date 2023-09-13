package services

import "main/models"

type CourseService interface {
	GetCourses() (*[]models.Course, error)
	CreateCourse(*models.Course) error
	UpdateCourse(*string, *models.Course) error
	DeleteCourse(*string) error
	AddCourses(*[]models.Course) error
	GetCourse(*string) (*models.Course, error)
}
