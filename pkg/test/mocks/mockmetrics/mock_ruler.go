// Code generated by mockery. DO NOT EDIT.

package mockmetrics

import (
	model "github.com/grafana/pyroscope/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// MockRuler is an autogenerated mock type for the Ruler type
type MockRuler struct {
	mock.Mock
}

type MockRuler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRuler) EXPECT() *MockRuler_Expecter {
	return &MockRuler_Expecter{mock: &_m.Mock}
}

// RecordingRules provides a mock function with given fields: tenant
func (_m *MockRuler) RecordingRules(tenant string) []*model.RecordingRule {
	ret := _m.Called(tenant)

	if len(ret) == 0 {
		panic("no return value specified for RecordingRules")
	}

	var r0 []*model.RecordingRule
	if rf, ok := ret.Get(0).(func(string) []*model.RecordingRule); ok {
		r0 = rf(tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RecordingRule)
		}
	}

	return r0
}

// MockRuler_RecordingRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecordingRules'
type MockRuler_RecordingRules_Call struct {
	*mock.Call
}

// RecordingRules is a helper method to define mock.On call
//   - tenant string
func (_e *MockRuler_Expecter) RecordingRules(tenant interface{}) *MockRuler_RecordingRules_Call {
	return &MockRuler_RecordingRules_Call{Call: _e.mock.On("RecordingRules", tenant)}
}

func (_c *MockRuler_RecordingRules_Call) Run(run func(tenant string)) *MockRuler_RecordingRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRuler_RecordingRules_Call) Return(_a0 []*model.RecordingRule) *MockRuler_RecordingRules_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRuler_RecordingRules_Call) RunAndReturn(run func(string) []*model.RecordingRule) *MockRuler_RecordingRules_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRuler creates a new instance of MockRuler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRuler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRuler {
	mock := &MockRuler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
