package delivery

import (
	"articles/app/helper"
	"articles/features/articles"
	"articles/features/articles/delivery/request"
	"articles/features/articles/delivery/response"

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

	if res == "can't data input" {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseCreateSuccess(res))
}

func (h *PostHandler) Get(c echo.Context) error {
	res, err := h.postBusiness.GetPost()
	if err == echo.ErrInternalServerError {
		return c.JSON(helper.ResponseInternalServerError(err.Error()))
	}
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkWithData("get data articles success", response.FromCoreToList(res)))
}
