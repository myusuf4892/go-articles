package delivery

import (
	"articles/app/helper"
	"articles/features/categories"
	"articles/features/categories/delivery/request"
	"articles/features/categories/delivery/response"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryBusiness categories.Business
}

func NewCategoryHandler(dataCtgy categories.Business) *CategoryHandler {
	return &CategoryHandler{
		categoryBusiness: dataCtgy,
	}
}

func (h *CategoryHandler) Create(c echo.Context) error {
	dataReq := request.Category{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("check your input, request input failed"))
	}

	dataCore := request.RequestToCore(dataReq)

	res, err := h.categoryBusiness.AddCategory(dataCore)

	if res == "400" {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseCreateSuccess("categories input success"))
}

func (h *CategoryHandler) Get(c echo.Context) error {
	res, err := h.categoryBusiness.GetCategory()
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkWithData("get data categories success", response.FromCoreToListResponse(res)))
}
