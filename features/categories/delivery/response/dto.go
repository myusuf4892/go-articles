package response

import "articles/features/categories"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCoreToResponse(data categories.Core) Category {
	return Category{
		ID:   data.ID,
		Name: data.Name,
	}
}

func FromCoreToListResponse(data []categories.Core) []Category {
	result := []Category{}
	for key := range data {
		result = append(result, FromCoreToResponse(data[key]))
	}
	return result
}
