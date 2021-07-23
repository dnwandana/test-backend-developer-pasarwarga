package service

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
	"github.com/dnwandana/test-backend-developer-pasarwarga/repository"
)

type ArticleServiceImpl struct {
	articleRepository *repository.ArticleRepository
}

func NewArticleService(repo *repository.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{
		articleRepository: repo,
	}
}

func (service *ArticleServiceImpl) Create(request *entity.Article) {
	panic("implement me")
}

func (service *ArticleServiceImpl) List() *[]model.ArticleResponse {
	panic("implement me")
}

func (service *ArticleServiceImpl) ListSoftDeleted() *[]model.ArticleResponse {
	panic("implement me")
}

func (service *ArticleServiceImpl) FindOne(articleID int64) {
	panic("implement me")
}

func (service *ArticleServiceImpl) Update(articleID int64, request *entity.Category) {
	panic("implement me")
}

func (service *ArticleServiceImpl) SoftDelete(articleID int64) {
	panic("implement me")
}

func (service *ArticleServiceImpl) Delete(articleID int64) {
	panic("implement me")
}
