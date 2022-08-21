package business

import (
	"articles/features/articles"
	"articles/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddPostSuccess(t *testing.T) {
	repo := new(mocks.PostData)

	dataReq := articles.Core{
		ID:         1,
		Title:      "Sample article 1",
		CategoryID: 1,
	}

	t.Run("Test Posting Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()
		srv := NewArticleBusiness(repo)

		row, err := srv.AddPost(dataReq)

		assert.Nil(t, err)
		assert.Equal(t, "201", row)
		repo.AssertExpectations(t)
	})
}

func TestAddPostFailed(t *testing.T) {
	repo := new(mocks.PostData)

	dataReq := articles.Core{
		ID:         0,
		Title:      "Sample article 1",
		CategoryID: 0,
	}

	t.Run("Test Posting Failed", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(0, errors.New("error server"))
		srv := NewArticleBusiness(repo)

		row, err := srv.AddPost(dataReq)

		assert.NotNil(t, err)
		assert.Equal(t, "400", row)
	})
}

func TestGetPostSuccess(t *testing.T) {
	repo := new(mocks.PostData)

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
	}

	t.Run("Test Get Posting Success", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(returnData, nil).Once()
		srv := NewArticleBusiness(repo)

		res, err := srv.GetPost()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0], res[0])
		repo.AssertExpectations(t)
	})
}

func TestGetPostFailed(t *testing.T) {
	repo := new(mocks.PostData)

	t.Run("Test Get Posting Failed", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]articles.Core{}, errors.New("error server"))
	})
	srv := NewArticleBusiness(repo)

	res, err := srv.GetPost()
	assert.NotNil(t, err)
	assert.Equal(t, []articles.Core{}, res)
}
