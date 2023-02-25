package repository

import "golang_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
