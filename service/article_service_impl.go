package service

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
	"github.com/dnwandana/test-backend-developer-pasarwarga/repository"
	"github.com/dnwandana/test-backend-developer-pasarwarga/util"
	"github.com/gosimple/slug"
	"strconv"
)

type ArticleServiceImpl struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(repo *repository.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{
		articleRepository: *repo,
	}
}

func (service *ArticleServiceImpl) Create(request *model.ArticleCreateRequest) {
	articleSlug := slug.Make(request.Title)
	article := entity.Article{
		Title:      request.Title,
		Slug:       articleSlug,
		CategoryID: request.CategoryID,
		Content:    request.Content,
	}
	_, txErr := service.articleRepository.Insert(&article)
	util.ReturnErrorIfNeeded(txErr)
}

func (service *ArticleServiceImpl) List() *[]model.ArticleResponse {
	articles, txErr := service.articleRepository.FindAll()
	util.ReturnErrorIfNeeded(txErr)
	return articles
}

func (service *ArticleServiceImpl) ListSoftDeleted() *[]model.ArticleResponse {
	articles, txErr := service.articleRepository.FindAllSoftDeleted()
	util.ReturnErrorIfNeeded(txErr)
	return articles
}

func (service *ArticleServiceImpl) FindOne(articleID string) *model.ArticleResponse {
	id, err := strconv.Atoi(articleID)
	util.ReturnErrorIfNeeded(err)

	article, txErr := service.articleRepository.FindByID(int64(id))
	util.ReturnErrorIfNeeded(txErr)

	return article
}

func (service *ArticleServiceImpl) Update(articleID string, request *model.ArticleUpdateRequest) {
	id, err := strconv.Atoi(articleID)
	util.ReturnErrorIfNeeded(err)

	articleSlug := slug.Make(request.Title)

	article := entity.Article{
		Title:      request.Title,
		Slug:       articleSlug,
		CategoryID: request.CategoryID,
		Content:    request.Content,
	}

	_, txErr := service.articleRepository.Update(int64(id), &article)
	util.ReturnErrorIfNeeded(txErr)
}

func (service *ArticleServiceImpl) SoftDelete(articleID string) {
	id, err := strconv.Atoi(articleID)
	util.ReturnErrorIfNeeded(err)

	_, txErr := service.articleRepository.SoftDelete(int64(id))
	util.ReturnErrorIfNeeded(txErr)
}

func (service *ArticleServiceImpl) Delete(articleID string) {
	id, err := strconv.Atoi(articleID)
	util.ReturnErrorIfNeeded(err)

	_, txErr := service.articleRepository.Delete(int64(id))
	util.ReturnErrorIfNeeded(txErr)
}
