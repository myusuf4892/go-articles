package business

import (
	"articles/features/categories"
	"articles/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddCategory(t *testing.T) {
	repo := new(mocks.CategoryData)

	dataReq := categories.Core{
		Name: "success-story",
	}

	t.Run("Test Add Category Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()
		srv := NewCategoryBusiness(repo)

		row, err := srv.AddCtgy(dataReq)

		assert.Nil(t, err)
		assert.Equal(t, "201", row)
		repo.AssertExpectations(t)
	})

	t.Run("Test Add Category Failed", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(0, errors.New("error input category")).Once()
		srv := NewCategoryBusiness(repo)

		row, err := srv.AddCtgy(dataReq)

		assert.NotNil(t, err)
		assert.Equal(t, "400", row)
	})
}

func TestGetCategory(t *testing.T) {
	repo := new(mocks.CategoryData)

	returnData := []categories.Core{
		{
			ID:   1,
			Name: "success-story",
		},
	}

	t.Run("Test Get Category Success", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(returnData, nil).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.GetCtgy()

		assert.Nil(t, err)
		assert.Equal(t, returnData[0], res[0])
		repo.AssertExpectations(t)
	})

	t.Run("Test Get Category Failed", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]categories.Core{}, errors.New("error server"))
		srv := NewCategoryBusiness(repo)

		res, err := srv.GetCtgy()

		assert.NotNil(t, err)
		assert.Equal(t, []categories.Core{}, res)
	})
}
