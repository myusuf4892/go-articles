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
		return c.JSON(helper.ResponseBadRequest("check your input, request input failed"))
	}

	dataCore := request.ToCore(dataReq)

	res, err := h.postBusiness.AddPost(dataCore)

	if res == "400" {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseCreateSuccess("article post success"))
}

func (h *PostHandler) Get(c echo.Context) error {
	res, err := h.postBusiness.GetPost()

	if err != nil {
		return c.JSON(helper.ResponseBadRequest("Failed Get all Posting"))
	}
	return c.JSON(helper.ResponseStatusOkWithData("get data articles success", response.FromCoreToList(res)))
}
