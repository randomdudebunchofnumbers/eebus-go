// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// TimeSeriesClientInterface is an autogenerated mock type for the TimeSeriesClientInterface type
type TimeSeriesClientInterface struct {
	mock.Mock
}

type TimeSeriesClientInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *TimeSeriesClientInterface) EXPECT() *TimeSeriesClientInterface_Expecter {
	return &TimeSeriesClientInterface_Expecter{mock: &_m.Mock}
}

// RequestConstraints provides a mock function with given fields:
func (_m *TimeSeriesClientInterface) RequestConstraints() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestConstraints")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.MsgCounterType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.MsgCounterType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeSeriesClientInterface_RequestConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestConstraints'
type TimeSeriesClientInterface_RequestConstraints_Call struct {
	*mock.Call
}

// RequestConstraints is a helper method to define mock.On call
func (_e *TimeSeriesClientInterface_Expecter) RequestConstraints() *TimeSeriesClientInterface_RequestConstraints_Call {
	return &TimeSeriesClientInterface_RequestConstraints_Call{Call: _e.mock.On("RequestConstraints")}
}

func (_c *TimeSeriesClientInterface_RequestConstraints_Call) Run(run func()) *TimeSeriesClientInterface_RequestConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TimeSeriesClientInterface_RequestConstraints_Call) Return(_a0 *model.MsgCounterType, _a1 error) *TimeSeriesClientInterface_RequestConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TimeSeriesClientInterface_RequestConstraints_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *TimeSeriesClientInterface_RequestConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// RequestData provides a mock function with given fields:
func (_m *TimeSeriesClientInterface) RequestData() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestData")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.MsgCounterType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.MsgCounterType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeSeriesClientInterface_RequestData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestData'
type TimeSeriesClientInterface_RequestData_Call struct {
	*mock.Call
}

// RequestData is a helper method to define mock.On call
func (_e *TimeSeriesClientInterface_Expecter) RequestData() *TimeSeriesClientInterface_RequestData_Call {
	return &TimeSeriesClientInterface_RequestData_Call{Call: _e.mock.On("RequestData")}
}

func (_c *TimeSeriesClientInterface_RequestData_Call) Run(run func()) *TimeSeriesClientInterface_RequestData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TimeSeriesClientInterface_RequestData_Call) Return(_a0 *model.MsgCounterType, _a1 error) *TimeSeriesClientInterface_RequestData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TimeSeriesClientInterface_RequestData_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *TimeSeriesClientInterface_RequestData_Call {
	_c.Call.Return(run)
	return _c
}

// RequestDescriptions provides a mock function with given fields:
func (_m *TimeSeriesClientInterface) RequestDescriptions() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestDescriptions")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.MsgCounterType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.MsgCounterType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeSeriesClientInterface_RequestDescriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestDescriptions'
type TimeSeriesClientInterface_RequestDescriptions_Call struct {
	*mock.Call
}

// RequestDescriptions is a helper method to define mock.On call
func (_e *TimeSeriesClientInterface_Expecter) RequestDescriptions() *TimeSeriesClientInterface_RequestDescriptions_Call {
	return &TimeSeriesClientInterface_RequestDescriptions_Call{Call: _e.mock.On("RequestDescriptions")}
}

func (_c *TimeSeriesClientInterface_RequestDescriptions_Call) Run(run func()) *TimeSeriesClientInterface_RequestDescriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TimeSeriesClientInterface_RequestDescriptions_Call) Return(_a0 *model.MsgCounterType, _a1 error) *TimeSeriesClientInterface_RequestDescriptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TimeSeriesClientInterface_RequestDescriptions_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *TimeSeriesClientInterface_RequestDescriptions_Call {
	_c.Call.Return(run)
	return _c
}

// WriteData provides a mock function with given fields: data
func (_m *TimeSeriesClientInterface) WriteData(data []model.TimeSeriesDataType) (*model.MsgCounterType, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for WriteData")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func([]model.TimeSeriesDataType) (*model.MsgCounterType, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func([]model.TimeSeriesDataType) *model.MsgCounterType); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func([]model.TimeSeriesDataType) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeSeriesClientInterface_WriteData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteData'
type TimeSeriesClientInterface_WriteData_Call struct {
	*mock.Call
}

// WriteData is a helper method to define mock.On call
//   - data []model.TimeSeriesDataType
func (_e *TimeSeriesClientInterface_Expecter) WriteData(data interface{}) *TimeSeriesClientInterface_WriteData_Call {
	return &TimeSeriesClientInterface_WriteData_Call{Call: _e.mock.On("WriteData", data)}
}

func (_c *TimeSeriesClientInterface_WriteData_Call) Run(run func(data []model.TimeSeriesDataType)) *TimeSeriesClientInterface_WriteData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]model.TimeSeriesDataType))
	})
	return _c
}

func (_c *TimeSeriesClientInterface_WriteData_Call) Return(_a0 *model.MsgCounterType, _a1 error) *TimeSeriesClientInterface_WriteData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TimeSeriesClientInterface_WriteData_Call) RunAndReturn(run func([]model.TimeSeriesDataType) (*model.MsgCounterType, error)) *TimeSeriesClientInterface_WriteData_Call {
	_c.Call.Return(run)
	return _c
}

// NewTimeSeriesClientInterface creates a new instance of TimeSeriesClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTimeSeriesClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TimeSeriesClientInterface {
	mock := &TimeSeriesClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}