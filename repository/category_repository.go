package repository

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
)

type CategoryRepository interface {
	Insert(request *entity.Category) (bool, error)

	FindAll() (*[]model.CategoryResponse, error)

	FindAllSoftDeleted() (*[]model.CategoryResponse, error)

	FindByID(categoryID int64) (*model.CategoryResponse, error)

	Update(categoryID int64, request *entity.Category) (bool, error)

	SoftDelete(categoryID int64) (bool, error)

	Delete(categoryID int64) (bool, error)
}
