package response

import (
	"articles/features/articles"
	"testing"
)

func TestResponseArticle(t *testing.T) {
	dataRes := articles.Core{
		ID:         1,
		Title:      "Sample article 1",
		CategoryID: 1,
		Category: articles.Category{
			ID:   1,
			Name: "success-story",
		},
	}

	FromCore(dataRes)
}

func TestResponseListArticles(t *testing.T) {
	dataRes := []articles.Core{
		{
			ID:         1,
			Title:      "Sample article 1",
			CategoryID: 1,
			Category: articles.Category{
				ID:   1,
				Name: "success-story",
			},
		},
		{
			ID:         2,
			Title:      "Sample article 2",
			CategoryID: 1,
			Category: articles.Category{
				ID:   1,
				Name: "success-story",
			},
		},
	}

	FromCoreToList(dataRes)
}
