package business

import "articles/features/articles"

type postUseCase struct {
	postData articles.Data
}

func NewArticleBusiness(dataPost articles.Data) articles.Business {
	return &postUseCase{
		postData: dataPost,
	}
}

func (uc *postUseCase) AddPost(dataReq articles.Core) (res string, err error) {
	row, err := uc.postData.Insert(dataReq)
	if row == 0 {
		res = "can't data input"
		return res, err
	}

	res = "article post success"
	return res, nil
}

func (uc *postUseCase) GetPost() (res []articles.Core, err error) {
	res, err = uc.postData.Get()
	if err != nil {
		return []articles.Core{}, err
	}
	return res, nil
}
