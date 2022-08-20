package mocks

import (
	"articles/features/articles"

	"github.com/stretchr/testify/mock"
)

type PostUseCase struct {
	mock.Mock
}

func (m *PostUseCase) AddPost(dataReq articles.Core) (res string, err error) {
	ret := m.Called(dataReq)

	var r0 int
	if rf, ok := ret.Get(0).(func(articles.Core) int); ok {
		r0 = rf(dataReq)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(articles.Core) error); ok {
		r1 = rf(dataReq)
	} else {
		r1 = ret.Error(1)
	}

	if r0 == 0 {
		res = "can't data input"
	}
	res = "success post articles"

	return res, r1
}

func (m *PostUseCase) GetPost() (dataPost []articles.Core, err error) {
	ret := m.Called()

	var r0 []articles.Core

	if rf, ok := ret.Get(0).(func() []articles.Core); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]articles.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
