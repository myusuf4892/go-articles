package mocks

import (
	"articles/features/categories"

	"github.com/stretchr/testify/mock"
)

type CtgyUseCase struct {
	mock.Mock
}

func (m *CtgyUseCase) AddCtgy(dataReq categories.Core) (res string, err error) {
	ret := m.Called(dataReq)

	var r0 string
	if rf, ok := ret.Get(0).(func(categories.Core) string); ok {
		r0 = rf(dataReq)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(categories.Core) error); ok {
		r1 = rf(dataReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *CtgyUseCase) GetCtgy() (dataPost []categories.Core, err error) {
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
