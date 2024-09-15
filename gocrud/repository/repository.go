package repository

import (
	"gocrud/model"
)

type UsersRepository interface {
	Save(tags model.Users)
	Update(tags model.Users)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Users, err error)
	FindAll() []model.Users
}
