package data

import (
	"articles/config"
	"articles/features/categories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbTest = config.InitDBTest()

func TestInsertSuccess(t *testing.T) {
	dbTest.Migrator().DropTable(&Category{})
	dbTest.AutoMigrate(&Category{})

	repo := NewCategoryRepo(dbTest)

	t.Run("Test Insert Category Success", func(t *testing.T) {
		mockCategory := categories.Core{
			Name: "success-story",
		}
		row, err := repo.Insert(mockCategory)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
}

func TestInsertFailed(t *testing.T) {
	dbTest.AutoMigrate(&Category{})

	repo := NewCategoryRepo(dbTest)

	t.Run("Test Insert Category Failed", func(t *testing.T) {
		mockCategory := categories.Core{
			Name: "success-story",
		}

		row, err := repo.Insert(mockCategory)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestGetSuccess(t *testing.T) {
	dbTest.AutoMigrate(&Category{})

	mockCategory := categories.Core{
		Name: "success-story",
	}

	repo := NewCategoryRepo(dbTest)

	t.Run("Test Get Category", func(t *testing.T) {
		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.Name, res[0].Name)
	})
}

func TestGetFailed(t *testing.T) {
	dbTest.Migrator().DropTable(&Category{})

	repo := NewCategoryRepo(dbTest)

	t.Run("Test Get Category", func(t *testing.T) {
		res, err := repo.Get()
		assert.NotNil(t, err)
		assert.Equal(t, []categories.Core{}, res)
	})
}
