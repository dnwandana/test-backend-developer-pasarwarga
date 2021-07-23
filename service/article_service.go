package service

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
)

type ArticleService interface {
	Create(request *model.ArticleCreateRequest)

	List() *[]model.ArticleResponse

	ListSoftDeleted() *[]model.ArticleResponse

	FindOne(articleID string) *model.ArticleResponse

	Update(articleID string, request *model.ArticleUpdateRequest)

	SoftDelete(articleID string)

	Delete(articleID string)
}
