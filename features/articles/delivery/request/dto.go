package request

import "articles/features/articles"

type Article struct {
	Title      string `json:"title" form:"title"`
	CategoryID int    `json:"category_id" form:"category_id"`
}

func ToCore(dataReq Article) articles.Core {
	core := articles.Core{
		Title:      dataReq.Title,
		CategoryID: dataReq.CategoryID,
	}
	return core
}
