package request

import "articles/features/categories"

type Category struct {
	Name string `json:"name" form:"name"`
}

func ToCore(dataReq Category) categories.Core {
	core := categories.Core{
		Name: dataReq.Name,
	}
	return core
}
