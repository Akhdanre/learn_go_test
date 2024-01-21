package repository

import "learn_go/learn_go_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}