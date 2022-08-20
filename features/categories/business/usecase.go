package business

import (
	"articles/features/categories"
)

type ctgyUseCase struct {
	ctgyData categories.Data
}

func NewCategoryBusiness(dataCtgy categories.Data) categories.Business {
	return &ctgyUseCase{
		ctgyData: dataCtgy,
	}
}

func (uc *ctgyUseCase) AddCtgy(dataReq categories.Core) (res string, err error) {
	row, err := uc.ctgyData.Insert(dataReq)
	if row == 0 {
		res = "can't data input"
		return res, err
	}

	res = "categories input success"
	return res, nil
}

func (uc *ctgyUseCase) GetCtgy() (res []categories.Core, err error) {
	res, err = uc.ctgyData.Get()
	if err != nil {
		return []categories.Core{}, err
	}
	return res, nil
}
