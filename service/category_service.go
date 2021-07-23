package service

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
)

type CategoryService interface {
	Create(request *model.CategoryCreateRequest)

	List() *[]model.CategoryResponse

	ListSoftDeleted() *[]model.CategoryResponse

	FindOne(categoryID string) *model.CategoryResponse

	Update(categoryID string, request *model.CategoryUpdateRequest)

	SoftDelete(categoryID string)

	Delete(categoryID string)
}
