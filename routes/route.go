package routes

import (
	factory "articles/app"
	"articles/app/helper"

	"github.com/labstack/echo/v4"
)

func New(present factory.Presenter) *echo.Echo {
	e := echo.New()
	e.Pre(helper.RemoveTrailingSlash())

	e.Use(helper.CorsMiddleware())

	e.POST("/articles", present.ArticlePresenter.Create)
	return e
}
