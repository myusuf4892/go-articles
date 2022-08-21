package data

import (
	"articles/features/categories"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (ctgy *Category) toCore() categories.Core {
	return categories.Core{
		ID:   int(ctgy.ID),
		Name: ctgy.Name,
	}
}

func toCoreList(data []Category) []categories.Core {
	result := []categories.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(ctgy *categories.Core) Category {
	return Category{
		Name: ctgy.Name,
	}
}
