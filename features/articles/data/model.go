package data

import (
	"articles/features/articles"
	"articles/features/categories/data"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string
	CategoryID int
	Category   data.Category
}

func (post *Article) toCore() articles.Core {
	return articles.Core{
		ID:         int(post.ID),
		Title:      post.Title,
		CategoryID: post.CategoryID,
		Category: articles.Category{
			ID:   int(post.Category.ID),
			Name: post.Category.Name,
		},
	}
}

func toCoreList(data []Article) []articles.Core {
	result := []articles.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(post *articles.Core) Article {
	return Article{
		Title:      post.Title,
		CategoryID: post.CategoryID,
		Category: data.Category{
			Name: post.Category.Name,
		},
	}
}
