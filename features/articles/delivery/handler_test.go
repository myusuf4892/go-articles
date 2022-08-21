package delivery

import (
	"articles/features/articles"
	"articles/features/articles/delivery/request"
	"articles/features/articles/delivery/response"
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

type AriclesResponseSuccess struct {
	Message string
	Data    []response.Article
}

type ResponseGlobal struct {
	Message string
}

var (
	mockData = request.Article{
		Title:      "Sample article 1",
		CategoryID: 1,
	}
)

var (
	mockDataFail = request.Article{
		Title:      "Sample article 1",
		CategoryID: 0,
	}
)

func TestCreatePost(t *testing.T) {
	dataReq, err := json.Marshal(mockData)
	if err != nil {
		t.Error(t, err, "error")
	}
	dataReqFail, err := json.Marshal(mockDataFail)
	if err != nil {
		t.Error(t, err, "error")
	}

	e := echo.New()
	usecase := new(mocks.PostUseCase)

	t.Run("Test Create Post Success", func(t *testing.T) {
		usecase.On("AddPost", mock.Anything).Return("201", nil).Once()

		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(dataReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Create(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "article post success", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Test Create Post Failed", func(t *testing.T) {
		usecase.On("AddPost", mock.Anything).Return("400", errors.New("error add posting")).Once()

		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(dataReqFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

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

	t.Run("Test Create Post Error Data Input", func(t *testing.T) {
		var dataReqFail = map[string]int{
			"title": 1,
		}
		reqBodyFail, _ := json.Marshal(dataReqFail)

		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

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
		usecase.AssertExpectations(t)
	})
}

func TestGetPosting(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.PostUseCase)

	returnData := []articles.Core{
		{
			ID:         1,
			Title:      "Sample article 1",
			CategoryID: 1,
			Category: articles.Category{
				ID:   1,
				Name: "success-story",
			},
		},
		{
			ID:         2,
			Title:      "Sample article 2",
			CategoryID: 1,
			Category: articles.Category{
				ID:   1,
				Name: "success-story",
			},
		},
	}

	returnDataFail := []articles.Core{}

	t.Run("Test Get data Posting Success", func(t *testing.T) {
		usecase.On("GetPost", mock.Anything).Return(returnData, nil).Once()
		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		responseData := AriclesResponseSuccess{}

		if assert.NoError(t, srv.Get(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0].Title, responseData.Data[0].Title)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Test Get data Posting Failed", func(t *testing.T) {
		usecase.On("GetPost", mock.Anything).Return(returnDataFail, errors.New("Failed Get all Posting")).Once()
		srv := NewArticleHandler(usecase)

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
			assert.Equal(t, "Failed Get all Posting", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}
