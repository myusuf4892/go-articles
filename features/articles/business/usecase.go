package business

import "articles/features/articles"

type articleUseCase struct {
	postData articles.Data
}

func NewArticleBusiness(dataPost articles.Data) articles.Business {
	return &articleUseCase{
		postData: dataPost,
	}
}

func (uc *articleUseCase) AddPost(dataReq articles.Core) (res string, err error) {
	row, err := uc.postData.Insert(dataReq)
	if err != nil {
		res = "error server"
		return res, err
	}
	if row == 0 {
		res = "can't data input"
		return res, err
	}

	res = "article post success"
	return res, nil
}

func (uc *articleUseCase) GetPost() (res []articles.Core, err error) {
	res, err = uc.postData.Get()
	if err != nil {
		return nil, err
	}
	return res, nil
}
