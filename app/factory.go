package app

import (
	postUseCase "articles/features/articles/business"
	postRepo "articles/features/articles/data"
	postHandler "articles/features/articles/delivery"

	ctgyUseCase "articles/features/categories/business"
	ctgyRepo "articles/features/categories/data"
	ctgyHandler "articles/features/categories/delivery"

	"gorm.io/gorm"
)

type Presenter struct {
	ArticlePresenter  *postHandler.PostHandler
	CategoryPresenter *ctgyHandler.CtgyHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	postData := postRepo.NewArticleRepo(dbConn)
	postBusiness := postUseCase.NewArticleBusiness(postData)
	postDelivery := postHandler.NewArticleHandler(postBusiness)

	ctgyData := ctgyRepo.NewCategoryRepo(dbConn)
	ctgyBusiness := ctgyUseCase.NewCategoryBusiness(ctgyData)
	ctgyDelivery := ctgyHandler.NewCategoryHandler(ctgyBusiness)

	return Presenter{
		ArticlePresenter:  postDelivery,
		CategoryPresenter: ctgyDelivery,
	}
}
