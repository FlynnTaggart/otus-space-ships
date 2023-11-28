// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UObject is an autogenerated mock type for the UObject type
type UObject struct {
	mock.Mock
}

type UObject_Expecter struct {
	mock *mock.Mock
}

func (_m *UObject) EXPECT() *UObject_Expecter {
	return &UObject_Expecter{mock: &_m.Mock}
}

// GetProperty provides a mock function with given fields: key
func (_m *UObject) GetProperty(key string) (interface{}, error) {
	ret := _m.Called(key)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UObject_GetProperty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProperty'
type UObject_GetProperty_Call struct {
	*mock.Call
}

// GetProperty is a helper method to define mock.On call
//   - key string
func (_e *UObject_Expecter) GetProperty(key interface{}) *UObject_GetProperty_Call {
	return &UObject_GetProperty_Call{Call: _e.mock.On("GetProperty", key)}
}

func (_c *UObject_GetProperty_Call) Run(run func(key string)) *UObject_GetProperty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UObject_GetProperty_Call) Return(_a0 interface{}, _a1 error) *UObject_GetProperty_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UObject_GetProperty_Call) RunAndReturn(run func(string) (interface{}, error)) *UObject_GetProperty_Call {
	_c.Call.Return(run)
	return _c
}

// SetProperty provides a mock function with given fields: key, value
func (_m *UObject) SetProperty(key string, value interface{}) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UObject_SetProperty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetProperty'
type UObject_SetProperty_Call struct {
	*mock.Call
}

// SetProperty is a helper method to define mock.On call
//   - key string
//   - value interface{}
func (_e *UObject_Expecter) SetProperty(key interface{}, value interface{}) *UObject_SetProperty_Call {
	return &UObject_SetProperty_Call{Call: _e.mock.On("SetProperty", key, value)}
}

func (_c *UObject_SetProperty_Call) Run(run func(key string, value interface{})) *UObject_SetProperty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *UObject_SetProperty_Call) Return(_a0 error) *UObject_SetProperty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UObject_SetProperty_Call) RunAndReturn(run func(string, interface{}) error) *UObject_SetProperty_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUObject interface {
	mock.TestingT
	Cleanup(func())
}

// NewUObject creates a new instance of UObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUObject(t mockConstructorTestingTNewUObject) *UObject {
	mock := &UObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
