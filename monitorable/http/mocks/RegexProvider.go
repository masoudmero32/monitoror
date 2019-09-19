// Code generated by mockery v1.0.0. DO NOT EDIT.

// If you want to rebuild this file, make mock-monitorable

package mocks

import mock "github.com/stretchr/testify/mock"

import regexp "regexp"

// RegexProvider is an autogenerated mock type for the RegexProvider type
type RegexProvider struct {
	mock.Mock
}

// GetRegex provides a mock function with given fields:
func (_m *RegexProvider) GetRegex() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetRegexp provides a mock function with given fields:
func (_m *RegexProvider) GetRegexp() *regexp.Regexp {
	ret := _m.Called()

	var r0 *regexp.Regexp
	if rf, ok := ret.Get(0).(func() *regexp.Regexp); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*regexp.Regexp)
		}
	}

	return r0
}