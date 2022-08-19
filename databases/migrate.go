package databases

import (
	mArticle "articles/features/articles/data"
	mCategory "articles/features/categories/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(mCategory.Category{})
	db.AutoMigrate(mArticle.Article{})
}
