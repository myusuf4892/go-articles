package data

import (
	"articles/config"
	"articles/features/articles"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var dbTest = config.InitDBTest()

type Category struct {
	gorm.Model
	Name string
}

func TestInsertSuccess(t *testing.T) {
	dbTest.Migrator().DropTable(&Article{}, &Category{})
	dbTest.AutoMigrate(&Article{}, &Category{})

	repo := NewArticleRepo(dbTest)

	t.Run("Test Insert Posting Success", func(t *testing.T) {
		mockArticle := articles.Core{
			Title:      "Sample article 1",
			CategoryID: 1,
			Category: articles.Category{
				ID:   1,
				Name: "success-story",
			},
		}
		row, err := repo.Insert(mockArticle)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
}

func TestInsertFailed(t *testing.T) {
	dbTest.Migrator().DropTable(&Article{})

	repo := NewArticleRepo(dbTest)

	t.Run("Test Insert Posting Failed", func(t *testing.T) {
		mockArticle := articles.Core{
			Title:      "Sample article 1",
			CategoryID: 0,
		}
		row, err := repo.Insert(mockArticle)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestGetSuccess(t *testing.T) {
	dbTest.Migrator().DropTable(&Article{}, &Category{})
	dbTest.AutoMigrate(&Article{}, &Category{})

	mockArticle := articles.Core{
		ID:         1,
		Title:      "Sample article 1",
		CategoryID: 1,
		Category: articles.Category{
			ID:   1,
			Name: "success-story",
		},
	}

	data := fromCore(&mockArticle)

	dbTest.Save(&data)
	repo := NewArticleRepo(dbTest)

	t.Run("Test Get Posting Success", func(t *testing.T) {
		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockArticle, res[0])
	})

}

func TestGetFailed(t *testing.T) {
	dbTest.Migrator().DropTable(&Article{}, &Category{})

	repo := NewArticleRepo(dbTest)

	t.Run("Test Get Posting Failed", func(t *testing.T) {
		res, err := repo.Get()
		assert.NotNil(t, err)
		assert.Equal(t, []articles.Core{}, res)
	})
}
