package response

import "articles/features/categories"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data categories.Core) Category {
	return Category{
		ID:   data.ID,
		Name: data.Name,
	}
}

func FromCoreToList(data []categories.Core) []Category {
	result := []Category{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
