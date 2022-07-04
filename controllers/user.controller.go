package controllers

import "github.com/Sleeplessss/gin_mongo_driver/models"

type UserInterface interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}