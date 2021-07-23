package service

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
)

type ArticleService interface {
	Create(request *entity.Article)

	List() *[]model.ArticleResponse

	ListSoftDeleted() *[]model.ArticleResponse

	FindOne(articleID int64)

	Update(articleID int64, request *entity.Category)

	SoftDelete(articleID int64)

	Delete(articleID int64)
}