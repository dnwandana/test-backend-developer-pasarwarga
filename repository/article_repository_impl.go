package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/dnwandana/test-backend-developer-pasarwarga/model"
)

type ArticleRepositoryImpl struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &ArticleRepositoryImpl{
		DB: db,
	}
}

func (r *ArticleRepositoryImpl) Insert(request *entity.Article) (bool, error) {
	query := "INSERT INTO articles (title, slug, category_id, content) VALUES (?, ?, ?, ?)"
	_, txErr := r.DB.ExecContext(context.Background(), query, request.Title, request.Slug, request.CategoryID, request.Content)
	if txErr != nil {
		return false, txErr
	}

	return true, nil
}

func (r *ArticleRepositoryImpl) FindAll() (*[]model.ArticleResponse, error) {
	query := `SELECT a.id, a.title, a.slug, c.id AS category_id, c.category_name, c.category_slug, a.created_at, a.updated_at, a.deleted_at
				FROM articles AS a INNER JOIN categories AS c on a.category_id = c.id WHERE a.deleted_at IS NULL`
	rows, e1 := r.DB.QueryContext(context.Background(), query)
	if e1 != nil {
		return nil, e1
	}

	defer rows.Close()
	var articles []model.ArticleResponse
	if rows.Next() {
		article := model.ArticleResponse{}
		e2 := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Slug,
			&article.CategoryID,
			&article.CategoryName,
			&article.CategorySlug,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.DeletedAt,
		)

		if e2 != nil {
			return nil, e2
		}

		articles = append(articles, article)
		return &articles, nil
	}

	return nil, nil
}

func (r *ArticleRepositoryImpl) FindAllByTitle(title string) (*[]model.ArticleResponse, error) {
	query := "SELECT a.id, a.title, a.slug, c.id AS category_id, c.category_name, c.category_slug, a.created_at, a.updated_at, a.deleted_at FROM articles AS a INNER JOIN categories AS c on a.category_id = c.id WHERE a.title LIKE '%" + title + "%' AND a.deleted_at IS NULL"
	rows, e1 := r.DB.QueryContext(context.Background(), query)
	if e1 != nil {
		return nil, e1
	}

	defer rows.Close()
	var articles []model.ArticleResponse
	if rows.Next() {
		article := model.ArticleResponse{}
		e2 := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Slug,
			&article.CategoryID,
			&article.CategoryName,
			&article.CategorySlug,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.DeletedAt,
		)

		if e2 != nil {
			return nil, e2
		}

		articles = append(articles, article)
		return &articles, nil
	}

	return nil, nil
}

func (r *ArticleRepositoryImpl) FindAllSoftDeleted() (*[]model.ArticleResponse, error) {
	query := `SELECT a.id, a.title, a.slug, c.id AS category_id, c.category_name, c.category_slug, a.created_at, a.updated_at, a.deleted_at
				FROM articles AS a INNER JOIN categories AS c on a.category_id = c.id WHERE a.deleted_at IS NOT NULL`
	rows, e1 := r.DB.QueryContext(context.Background(), query)
	if e1 != nil {
		return nil, e1
	}

	defer rows.Close()
	var articles []model.ArticleResponse
	if rows.Next() {
		article := model.ArticleResponse{}
		e2 := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Slug,
			&article.CategoryID,
			&article.CategoryName,
			&article.CategorySlug,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.DeletedAt,
		)

		if e2 != nil {
			return nil, e2
		}

		articles = append(articles, article)
		return &articles, nil
	}

	return nil, nil
}

func (r *ArticleRepositoryImpl) FindByID(articleID int64) (*model.ArticleResponse, error) {
	query := `SELECT a.id, a.title, a.slug, c.id AS category_id, c.category_name, c.category_slug, a.created_at, a.updated_at, a.deleted_at
				FROM articles AS a INNER JOIN categories AS c on a.category_id = c.id WHERE a.id = ?`
	rows, e1 := r.DB.QueryContext(context.Background(), query, articleID)
	if e1 != nil {
		return nil, e1
	}

	defer rows.Close()
	if rows.Next() {
		article := model.ArticleResponse{}
		e2 := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Slug,
			&article.CategoryID,
			&article.CategoryName,
			&article.CategorySlug,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.DeletedAt,
		)

		if e2 != nil {
			return nil, e2
		}

		return &article, nil
	}

	return nil, nil
}

func (r *ArticleRepositoryImpl) Update(articleID int64, request *entity.Article) (bool, error) {
	query := "UPDATE articles SET title = ?, slug = ? , category_id = ?, content = ? WHERE id = ?"
	result, e1 := r.DB.ExecContext(context.Background(), query, request.Title, request.Slug, request.CategoryID, request.Content, articleID)
	if e1 != nil {
		return false, e1
	}

	affected, e2 := result.RowsAffected()
	if e2 != nil {
		return false, e2
	}

	if affected != 1 {
		return false, errors.New("no article updated")
	}

	return true, nil
}

func (r *ArticleRepositoryImpl) SoftDelete(articleID int64) (bool, error) {
	query := "UPDATE articles SET deleted_at = NOW() WHERE id = ?"
	result, e1 := r.DB.ExecContext(context.Background(), query, articleID)
	if e1 != nil {
		return false, e1
	}

	affected, e2 := result.RowsAffected()
	if e2 != nil {
		return false, e2
	}

	if affected != 1 {
		return false, errors.New("no article deleted")
	}

	return true, nil
}

func (r *ArticleRepositoryImpl) Delete(articleID int64) (bool, error) {
	query := "DELETE FROM articles WHERE id = ?"
	result, e1 := r.DB.ExecContext(context.Background(), query, articleID)
	if e1 != nil {
		return false, e1
	}

	affected, e2 := result.RowsAffected()
	if e2 != nil {
		return false, e2
	}

	if affected != 1 {
		return false, errors.New("no article deleted")
	}

	return true, nil
}
