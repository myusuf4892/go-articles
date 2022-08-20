package delivery

import (
	"articles/app/helper"
	"articles/features/categories"
	"articles/features/categories/delivery/request"
	"articles/features/categories/delivery/response"

	"github.com/labstack/echo/v4"
)

type CtgyHandler struct {
	ctgyBusiness categories.Business
}

func NewCategoryHandler(dataCtgy categories.Business) *CtgyHandler {
	return &CtgyHandler{
		ctgyBusiness: dataCtgy,
	}
}

func (h *CtgyHandler) Create(c echo.Context) error {
	dataReq := request.Category{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("request input failed"))
	}

	dataCore := request.ToCore(dataReq)

	res, err := h.ctgyBusiness.AddCtgy(dataCore)

	if res == "can't data input" {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseCreateSuccess(res))
}

func (h *CtgyHandler) Get(c echo.Context) error {
	res, err := h.ctgyBusiness.GetCtgy()
	if err == echo.ErrInternalServerError {
		return c.JSON(helper.ResponseInternalServerError(err.Error()))
	}
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkWithData("get data categories success", response.FromCoreToList(res)))
}
