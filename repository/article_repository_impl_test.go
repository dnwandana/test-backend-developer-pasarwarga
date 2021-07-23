package repository

import (
	"github.com/dnwandana/test-backend-developer-pasarwarga/database"
	"github.com/dnwandana/test-backend-developer-pasarwarga/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArticleRepositoryImpl_Insert(t *testing.T) {
	type args struct {
		request *entity.Article
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Success_1",
			args: args{request: &entity.Article{
				Title:      "Tutorial Optimize Query MySQL",
				Slug:       "tutorial-optimize-query-mysql",
				CategoryID: 1,
				Content:    "Tutorial Optimize Query MySQL",
			}},
			want:    true,
			wantErr: false,
		},
		{
			name: "Success_2",
			args: args{request: &entity.Article{
				Title:      "Tutorial Install Docker",
				Slug:       "tutorial-install-docker",
				CategoryID: 2,
				Content:    "Tutorial Install Docker",
			}},
			want:    true,
			wantErr: false,
		},
		{
			name: "Fail_1",
			args: args{request: &entity.Article{
				Title:      "Tutorial Optimize Query MySQL",
				Slug:       "tutorial-optimize-query-mysql",
				CategoryID: 1,
				Content:    "Tutorial Optimize Query MySQL",
			}},
			want:    false,
			wantErr: true,
		},
		{
			name: "Fail_2",
			args: args{request: &entity.Article{
				Title:      "Tutorial Install Docker",
				Slug:       "tutorial-install-docker",
				CategoryID: 2,
				Content:    "Tutorial Install Docker",
			}},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepositoryImpl{
				DB: database.GetConnection(),
			}
			got, err := r.Insert(tt.args.request)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestArticleRepositoryImpl_FindAll(t *testing.T) {
	result, txErr := NewArticleRepository(database.GetConnection()).FindAll()
	assert.Nil(t, txErr)
	assert.NotNil(t, result)
}

func TestArticleRepositoryImpl_FindAllSoftDeleted(t *testing.T) {
	result, txErr := NewArticleRepository(database.GetConnection()).FindAllSoftDeleted()
	assert.Nil(t, txErr)
	assert.NotNil(t, result)
}

func TestArticleRepositoryImpl_FindByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result, txErr := NewArticleRepository(database.GetConnection()).FindByID(1)
		assert.Nil(t, txErr)
		assert.NotNil(t, result)
	})

	t.Run("Fail", func(t *testing.T) {
		result, txErr := NewArticleRepository(database.GetConnection()).FindByID(0)
		assert.Error(t, txErr)
		assert.Nil(t, result)
	})
}

func TestArticleRepositoryImpl_Update(t *testing.T) {
	type args struct {
		articleID int64
		request   *entity.Article
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
				articleID: 1,
				request: &entity.Article{
					Title:      "Tutorial Optimize Query MySQL (Updated)",
					Slug:       "tutorial-optimize-query-mysql-updated",
					CategoryID: 1,
					Content:    "Tutorial Optimize Query MySQL (Updated)",
				},
			},
			want:    true,
			wantErr: false,
		}, {
			name: "Fail",
			args: args{
				articleID: 0,
				request: &entity.Article{
					Title:      "No Article",
					Slug:       "no-article",
					CategoryID: 0,
					Content:    "No Article",
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepositoryImpl{
				DB: database.GetConnection(),
			}
			got, err := r.Update(tt.args.articleID, tt.args.request)
			if (err != nil) != tt.wantErr {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestArticleRepositoryImpl_SoftDelete(t *testing.T) {
	type args struct {
		articleID int64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{articleID: 1},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Fail",
			args:    args{articleID: 0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepositoryImpl{
				DB: database.GetConnection(),
			}
			got, err := r.SoftDelete(tt.args.articleID)
			if (err != nil) != tt.wantErr {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestArticleRepositoryImpl_Delete(t *testing.T) {
	type args struct {
		articleID int64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{articleID: 1},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Fail",
			args:    args{articleID: 0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ArticleRepositoryImpl{
				DB: database.GetConnection(),
			}
			got, err := r.Delete(tt.args.articleID)
			if (err != nil) != tt.wantErr {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			if got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
