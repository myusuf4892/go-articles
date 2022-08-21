package mocks

import (
	"articles/features/categories"

	"github.com/stretchr/testify/mock"
)

type CategoryData struct {
	mock.Mock
}

func (m *CategoryData) Insert(dataReq categories.Core) (row int, err error) {
	ret := m.Called(dataReq)

	var r0 int
	if rf, ok := ret.Get(0).(func(categories.Core) int); ok {
		r0 = rf(dataReq)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(categories.Core) error); ok {
		r1 = rf(dataReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *CategoryData) Get() (dataPost []categories.Core, err error) {
	ret := m.Called()

	var r0 []categories.Core
	if rf, ok := ret.Get(0).(func() []categories.Core); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]categories.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
