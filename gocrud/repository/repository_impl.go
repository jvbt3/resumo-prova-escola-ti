package repository

import (
	"errors"
	"gocrud/data/request"
	"gocrud/helper"
	"gocrud/model"
	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

func (t UsersRepositoryImpl) Save(users model.Users) {
	result := t.Db.Create(&users)
	helper.ErrorPanic(result.Error)

}

func (t UsersRepositoryImpl) Update(users model.Users) {
	var updateTag = request.UpdateUsersRequest{
		Id:   users.Id,
		Name: users.Name,
	}
	result := t.Db.Model(&users).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func (t UsersRepositoryImpl) Delete(usersId int) {
	var users model.Users
	result := t.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

func (t UsersRepositoryImpl) FindById(usersId int) (model.Users, error) {
	var tag model.Users
	result := t.Db.Find(&tag, usersId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (t UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	results := t.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users
}
