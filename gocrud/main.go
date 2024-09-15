package main

import (
	"github.com/go-playground/validator/v10"
	"gocrud/config"
	"gocrud/controller"
	"gocrud/helper"
	"gocrud/model"
	"gocrud/repository"
	"gocrud/router"
	"gocrud/service"
	"net/http"
	"time"
)

func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)

	//Init Service
	userService := service.NewUserServiceImpl(userRepository, validate)

	//Init controller
	userController := controller.NewUserController(userService)

	//Router
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
