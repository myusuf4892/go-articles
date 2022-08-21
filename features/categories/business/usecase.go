package business

import (
	"articles/features/categories"
	"errors"
)

type categoryUseCase struct {
	categoryData categories.Data
}

func NewCategoryBusiness(dataCtgy categories.Data) categories.Business {
	return &categoryUseCase{
		categoryData: dataCtgy,
	}
}

func (uc *categoryUseCase) AddCategory(dataReq categories.Core) (res string, err error) {
	if dataReq.Name == "" {
		return "400", errors.New("must be filled")
	}
	row, err := uc.categoryData.Insert(dataReq)
	if row == 0 {
		res = "400"
		return res, err
	}

	res = "201"
	return res, nil
}

func (uc *categoryUseCase) GetCategory() (res []categories.Core, err error) {
	res, err = uc.categoryData.Get()
	if err != nil {
		return []categories.Core{}, err
	}
	return res, nil
}
