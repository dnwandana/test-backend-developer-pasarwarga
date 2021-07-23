package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (r *CategoryRepositoryImpl) Insert(request *entity.Category) (bool, error) {
	query := "INSERT INTO categories (category_name, category_slug) VALUES (?, ?)"
	_, txErr := r.DB.ExecContext(context.Background(), query, request.CategoryName, request.CategorySlug)
	if txErr != nil {
		return false, txErr
	}

	return true, nil
}

func (r *CategoryRepositoryImpl) FindAll() (*[]entity.Category, error) {
	query := "SELECT * FROM categories WHERE deleted_at IS NULL"
	rows, e1 := r.DB.QueryContext(context.Background(), query)
	if e1 != nil {
		return nil, e1
	}

	defer rows.Close()
	var categories []entity.Category
	if rows.Next() {
		category := entity.Category{}
		e2 := rows.Scan(
			&category.ID,
			&category.CategoryName,
			&category.CategorySlug,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.DeletedAt,
		)

		if e2 != nil {
			return nil, e2
		}

		categories = append(categories, category)
		return &categories, nil
	}

	return nil, nil
}

func (r *CategoryRepositoryImpl) FindAllSoftDeleted() (*[]entity.Category, error) {
	query := "SELECT * FROM categories WHERE deleted_at IS NOT NULL"
	rows, e1 := r.DB.QueryContext(context.Background(), query)
	if e1 != nil {
		return nil, e1
	}

	defer rows.Close()
	var categories []entity.Category
	if rows.Next() {
		category := entity.Category{}
		e2 := rows.Scan(
			&category.ID,
			&category.CategoryName,
			&category.CategorySlug,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.DeletedAt,
		)

		if e2 != nil {
			return nil, e2
		}

		categories = append(categories, category)
		return &categories, nil
	}

	return nil, nil
}

func (r *CategoryRepositoryImpl) FindByID(categoryID int64) (*entity.Category, error) {
	query := "SELECT * FROM categories WHERE id = ?"
	rows, e1 := r.DB.QueryContext(context.Background(), query, categoryID)
	if e1 != nil {
		return nil, e1
	}

	defer rows.Close()
	if rows.Next() {
		category := entity.Category{}
		e2 := rows.Scan(
			&category.ID,
			&category.CategoryName,
			&category.CategorySlug,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.DeletedAt,
		)

		if e2 != nil {
			return nil, e2
		}

		return &category, nil
	}

	return nil, nil
}

func (r *CategoryRepositoryImpl) Update(categoryID int64, request *entity.Category) (bool, error) {
	query := "UPDATE categories SET category_name = ?, category_slug = ? WHERE id = ?"
	result, e1 := r.DB.ExecContext(context.Background(), query, request.CategoryName, request.CategorySlug, categoryID)
	if e1 != nil {
		return false, e1
	}

	affected, e2 := result.RowsAffected()
	if e2 != nil {
		return false, e2
	}

	if affected != 1 {
		return false, errors.New("no category updated")
	}

	return true, nil
}

func (r *CategoryRepositoryImpl) SoftDelete(categoryID int64) (bool, error) {
	query := "UPDATE categories SET deleted_at = NOW() WHERE id = ?"
	result, e1 := r.DB.ExecContext(context.Background(), query, categoryID)
	if e1 != nil {
		return false, e1
	}

	affected, e2 := result.RowsAffected()
	if e2 != nil {
		return false, e2
	}

	if affected != 1 {
		return false, errors.New("no category deleted")
	}

	return true, nil
}

func (r *CategoryRepositoryImpl) Delete(categoryID int64) (bool, error) {
	query := "DELETE FROM categories WHERE id = ?"
	result, e1 := r.DB.ExecContext(context.Background(), query, categoryID)
	if e1 != nil {
		return false, e1
	}

	affected, e2 := result.RowsAffected()
	if e2 != nil {
		return false, e2
	}

	if affected != 1 {
		return false, errors.New("no category deleted")
	}

	return true, nil
}
