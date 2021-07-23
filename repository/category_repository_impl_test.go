package repository

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/database"
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryRepositoryImpl_Insert(t *testing.T) {
	type args struct {
		request *entity.Category
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Success_1",
			args: args{request: &entity.Category{
				CategoryName: "MySQL",
				CategorySlug: "mysql",
			}},
			want:    true,
			wantErr: false,
		},
		{
			name: "Success_2",
			args: args{request: &entity.Category{
				CategoryName: "Docker",
				CategorySlug: "docker",
			}},
			want:    true,
			wantErr: false,
		},
		{
			name: "Fail_1",
			args: args{request: &entity.Category{
				CategoryName: "MySQL",
				CategorySlug: "mysql",
			}},
			want:    false,
			wantErr: true,
		},
		{
			name: "Fail_2",
			args: args{request: &entity.Category{
				CategoryName: "Docker",
				CategorySlug: "docker",
			}},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepositoryImpl{
				DB: database.GetConnection(),
			}
			got, err := r.Insert(tt.args.request)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCategoryRepositoryImpl_FindAll(t *testing.T) {
	result, txErr := NewCategoryRepository(database.GetConnection()).FindAll()
	assert.Nil(t, txErr)
	assert.NotNil(t, result)
}

func TestCategoryRepositoryImpl_FindAllSoftDeleted(t *testing.T) {
	result, txErr := NewCategoryRepository(database.GetConnection()).FindAllSoftDeleted()
	assert.Nil(t, txErr)
	assert.Nil(t, result)
}

func TestCategoryRepositoryImpl_FindByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result, txErr := NewCategoryRepository(database.GetConnection()).FindByID(1)
		assert.Nil(t, txErr)
		assert.NotNil(t, result)
	})

	t.Run("Fail", func(t *testing.T) {
		result, txErr := NewCategoryRepository(database.GetConnection()).FindByID(0)
		assert.Error(t, txErr)
		assert.Nil(t, result)
	})
}

func TestCategoryRepositoryImpl_Update(t *testing.T) {
	type args struct {
		categoryID int64
		request    *entity.Category
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				categoryID: 1,
				request: &entity.Category{
					CategoryName: "MySQL Updated",
					CategorySlug: "mysql-updated",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				categoryID: 0,
				request: &entity.Category{
					CategoryName: "No Category Updated",
					CategorySlug: "no-category-updated",
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepositoryImpl{
				DB: database.GetConnection(),
			}
			got, err := r.Update(tt.args.categoryID, tt.args.request)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCategoryRepositoryImpl_SoftDelete(t *testing.T) {
	type args struct {
		categoryID int64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{categoryID: 1},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Fail",
			args:    args{categoryID: 0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepositoryImpl{
				DB: database.GetConnection(),
			}
			got, err := r.SoftDelete(tt.args.categoryID)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCategoryRepositoryImpl_Delete(t *testing.T) {
	type args struct {
		categoryID int64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{categoryID: 1},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Fail",
			args:    args{categoryID: 0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CategoryRepositoryImpl{
				DB: database.GetConnection(),
			}
			got, err := r.Delete(tt.args.categoryID)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
