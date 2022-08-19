package response

import "articles/features/articles"

type Article struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Category Category `json:"category"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data articles.Core) Article {
	return Article{
		ID:    data.ID,
		Title: data.Title,
		Category: Category{
			ID:   data.Category.ID,
			Name: data.Category.Name,
		},
	}
}

func FromCoreToList(data []articles.Core) []Article {
	result := []Article{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
