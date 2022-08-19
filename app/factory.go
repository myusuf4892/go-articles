package app

import (
	postUseCase "articles/features/articles/business"
	postRepo "articles/features/articles/data"
	"articles/features/articles/delivery"
	postHandler "articles/features/articles/delivery"

	"gorm.io/gorm"
)

type Presenter struct {
	ArticlePresenter *delivery.PostHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	postData := postRepo.NewArticleRepo(dbConn)
	postBusiness := postUseCase.NewArticleBusiness(postData)
	postDelivery := postHandler.NewArticleHandler(postBusiness)

	return Presenter{
		ArticlePresenter: postDelivery,
	}
}
