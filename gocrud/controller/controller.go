package controller

import (
	"github.com/gin-gonic/gin"
	"gocrud/data/request"
	"gocrud/data/response"
	"gocrud/helper"
	"gocrud/service"
	"net/http"
	"strconv"
)

type UserController struct {
	userService service.UsersService
}

func NewUserController(service service.UsersService) *UserController {
	return &UserController{userService: service}
}

func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.userService.Create(createUserRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUsersRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	updateUserRequest.Id = id

	controller.userService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.userService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.userService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	userResponse := controller.userService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
