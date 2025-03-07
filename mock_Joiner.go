// Code generated by mockery. DO NOT EDIT.

package patcher

import mock "github.com/stretchr/testify/mock"

// MockJoiner is an autogenerated mock type for the Joiner type
type MockJoiner struct {
	mock.Mock
}

// Join provides a mock function with no fields
func (_m *MockJoiner) Join() (string, []interface{}) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Join")
	}

	var r0 string
	var r1 []interface{}
	if rf, ok := ret.Get(0).(func() (string, []interface{})); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() []interface{}); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	return r0, r1
}

// NewMockJoiner creates a new instance of MockJoiner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockJoiner(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockJoiner {
	mock := &MockJoiner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
