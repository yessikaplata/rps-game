package mocks

import "github.com/stretchr/testify/mock"

type RandomGenerator struct {
	mock.Mock
}

func (_m *RandomGenerator) Intn(n int) int {
	ret := _m.Called(n)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewRandomGenerator interface {
	mock.TestingT
	Cleanup(func())
}

func NewRandomGenerator(t mockConstructorTestingTNewRandomGenerator) *RandomGenerator {
	mock := &RandomGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
