package repository

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
)

type CategoryRepository interface {
	Insert(request *entity.Category) (bool, error)

	FindAll() (*[]entity.Category, error)

	FindAllSoftDeleted() (*[]entity.Category, error)

	FindByID(categoryID int64) (*entity.Category, error)

	Update(categoryID int64, request *entity.Category) (bool, error)

	SoftDelete(categoryID int64) (bool, error)

	Delete(categoryID int64) (bool, error)
}
