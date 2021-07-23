package repository

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
)

type ArticleRepository interface {
	Insert(request *entity.Article) (bool, error)

	FindAll() (*[]model.ArticleResponse, error)

	FindAllSoftDeleted() (*[]model.ArticleResponse, error)

	FindByID(articleID int64) (*model.ArticleResponse, error)

	Update(articleID int64, request *entity.Article) (bool, error)

	SoftDelete(articleID int64) (bool, error)

	Delete(articleID int64) (bool, error)
}
