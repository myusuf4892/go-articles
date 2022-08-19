package delivery

import (
	"articles/app/helper"
	"articles/features/articles"
	"articles/features/articles/delivery/request"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postBusiness articles.Business
}

func NewArticleHandler(dataPost articles.Business) *PostHandler {
	return &PostHandler{
		postBusiness: dataPost,
	}
}

func (h *PostHandler) Create(c echo.Context) error {
	dataReq := request.Article{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("request input failed"))
	}

	dataCore := request.ToCore(dataReq)

	res, err := h.postBusiness.AddPost(dataCore)

	if res == "error server" {
		return c.JSON(helper.ResponseInternalServerError(err.Error()))
	}

	if res == "can't data input" {
		return c.JSON(helper.ResponseBadRequest(res))
	}

	return c.JSON(helper.ResponseCreateSuccess(res))
}
