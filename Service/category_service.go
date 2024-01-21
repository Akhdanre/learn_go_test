package service

import (
	"errors"
	"learn_go/learn_go_test/entity"
	"learn_go/learn_go_test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error){
	category := service.Repository.FindById(id)
	if category == nil{
		return category, errors.New("category not found")
	}else {
		return category, nil
	}
}