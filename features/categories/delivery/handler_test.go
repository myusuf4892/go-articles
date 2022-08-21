package delivery

import (
	"articles/features/articles/delivery/response"
	"articles/features/categories"
	"articles/features/categories/delivery/request"
	"articles/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CategoryResponseSuccess struct {
	Message string
	Data    []response.Category
}

type ResponseGlobal struct {
	Message string
}

var (
	mockData = request.Category{
		Name: "success-story",
	}
)

func TestCreateCategory(t *testing.T) {
	dataReq, err := json.Marshal(mockData)
	if err != nil {
		t.Error(t, err, "error")
	}

	e := echo.New()
	usecase := new(mocks.CategoryUseCase)

	t.Run("Test Create Category Success", func(t *testing.T) {
		usecase.On("AddCategory", mock.Anything).Return("201", nil).Once()

		srv := NewCategoryHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(dataReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/categories")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Create(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "categories input success", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Test Create Category Failed", func(t *testing.T) {
		usecase.On("AddCategory", mock.Anything).Return("400", errors.New("error create category"))

		srv := NewCategoryHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(dataReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/categories")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Create(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Test Create Error Data Input", func(t *testing.T) {
		var dataReqFail = map[string]int{
			"name": 1,
		}
		reqBodyFail, _ := json.Marshal(dataReqFail)

		srv := NewCategoryHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/categories")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Create(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "check your input, request input failed", responseData.Message)
		}
	})
}

func TestGetCategory(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.CategoryUseCase)

	returnData := []categories.Core{
		{
			ID:   1,
			Name: "success-story",
		},
	}

	returnDataFail := []categories.Core{}

	t.Run("Test Get Category Success", func(t *testing.T) {
		usecase.On("GetCategory", mock.Anything).Return(returnData, nil).Once()
		srv := NewCategoryHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		responseData := CategoryResponseSuccess{}

		if assert.NoError(t, srv.Get(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0].Name, responseData.Data[0].Name)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Test Get Category Failed", func(t *testing.T) {
		usecase.On("GetCategory", mock.Anything).Return(returnDataFail, errors.New("Failed Get all Category")).Once()
		srv := NewCategoryHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Get(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Failed Get all Category", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}
