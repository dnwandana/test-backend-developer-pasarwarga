package service

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
)

type CategoryService interface {
	Create(request *entity.Category)

	List() *[]model.CategoryResponse

	ListSoftDeleted() *[]model.CategoryResponse

	FindOne(categoryID int64)

	Update(categoryID int64, request *entity.Category)

	SoftDelete(categoryID int64)

	Delete(categoryID int64)
}
