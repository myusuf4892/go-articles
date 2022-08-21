package routes

import (
	factory "articles/app"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(present factory.Presenter) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	// routes articles
	e.POST("/articles", present.ArticlePresenter.Create)
	e.GET("/articles", present.ArticlePresenter.Get)
	// routes categories
	e.POST("/categories", present.CategoryPresenter.Create)
	e.GET("/categories", present.CategoryPresenter.Get)
	return e
}
