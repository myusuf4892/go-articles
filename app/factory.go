package app

import (
	postUseCase "articles/features/articles/business"
	postRepo "articles/features/articles/data"
	postHandler "articles/features/articles/delivery"

	categoryUseCase "articles/features/categories/business"
	categoryRepo "articles/features/categories/data"
	categoryHandler "articles/features/categories/delivery"

	"gorm.io/gorm"
)

type Presenter struct {
	ArticlePresenter  *postHandler.PostHandler
	CategoryPresenter *categoryHandler.CategoryHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	postData := postRepo.NewArticleRepo(dbConn)
	postBusiness := postUseCase.NewArticleBusiness(postData)
	postDelivery := postHandler.NewArticleHandler(postBusiness)

	categoryData := categoryRepo.NewCategoryRepo(dbConn)
	categoryBusiness := categoryUseCase.NewCategoryBusiness(categoryData)
	categoryDelivery := categoryHandler.NewCategoryHandler(categoryBusiness)

	return Presenter{
		ArticlePresenter:  postDelivery,
		CategoryPresenter: categoryDelivery,
	}
}
