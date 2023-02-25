package service

import (
	"golang_unit_test/entity"
	"golang_unit_test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var CategoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: CategoryRepository}

func TestcategoryService_Get(t *testing.T) {

	// program mock 
	CategoryRepository.Mock.On("FindById","").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T){
	category := entity.Category{
		Id: "1",
		Name: "Laptop",
	}
	CategoryRepository.Mock.On("FindById", "2").Return(category)

	result. err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}