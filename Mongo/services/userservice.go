package services

import "main/models"

type UserServices interface {
	GetUsers() (*[]models.User, error)
	CreateUser(*models.User) error
	DeleteUser(*string) error
	GetUser(*string) (*models.User, error)
	UpdateUser(*models.User) error
	CreateUsers(*[]models.User) error
}
