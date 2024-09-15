package service

import (
	"github.com/go-playground/validator/v10"
	"gocrud/data/request"
	"gocrud/data/response"
	"gocrud/helper"
	"gocrud/model"
	"gocrud/repository"
)

type UsersServiceImpl struct {
	UserRepository repository.UsersRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (t UsersServiceImpl) Create(user request.CreateUsersRequest) {
	err := t.Validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := model.Users{
		Id:   user.Id,
		Name: user.Name,
	}
	t.UserRepository.Save(userModel)
}

func (t UsersServiceImpl) Update(user request.UpdateUsersRequest) {
	userData, err := t.UserRepository.FindById(user.Id)
	helper.ErrorPanic(err)
	userData.Name = user.Name
	t.UserRepository.Update(userData)
}

func (t UsersServiceImpl) Delete(userId int) {
	t.UserRepository.Delete(userId)
}

func (t UsersServiceImpl) FindById(userId int) response.UsersResponse {
	userData, err := t.UserRepository.FindById(userId)
	helper.ErrorPanic(err)

	userResponse := response.UsersResponse{
		Id:   userData.Id,
		Name: userData.Name,
	}
	return userResponse
}

func (t UsersServiceImpl) FindAll() []response.UsersResponse {
	result := t.UserRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		users = append(users, user)
	}
	return users
}
