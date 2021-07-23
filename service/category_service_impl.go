package service

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
	"github.com/dnwandana/test-backend-developer-pasarwarga/repository"
)

type CategoryServiceImpl struct {
	categoryRepository *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		categoryRepository: repo,
	}
}

func (service *CategoryServiceImpl) Create(request *entity.Category) {
	panic("implement me")
}

func (service *CategoryServiceImpl) List() *[]model.CategoryResponse {
	panic("implement me")
}

func (service *CategoryServiceImpl) ListSoftDeleted() *[]model.CategoryResponse {
	panic("implement me")
}

func (service *CategoryServiceImpl) FindOne(categoryID int64) {
	panic("implement me")
}

func (service *CategoryServiceImpl) Update(categoryID int64, request *entity.Category) {
	panic("implement me")
}

func (service *CategoryServiceImpl) SoftDelete(categoryID int64) {
	panic("implement me")
}

func (service *CategoryServiceImpl) Delete(categoryID int64) {
	panic("implement me")
}
