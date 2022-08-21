package response

import (
	"articles/features/categories"
	"testing"
)

func TestResponseCategory(t *testing.T) {
	dataRes := categories.Core{
		ID:   1,
		Name: "success-story",
	}

	FromCoreToResponse(dataRes)
}

func TestResponseListCategories(t *testing.T) {
	dataRes := []categories.Core{
		{
			ID:   1,
			Name: "success-story",
		},
		{
			ID:   2,
			Name: "success-story-dummy",
		},
	}

	FromCoreToListResponse(dataRes)
}
