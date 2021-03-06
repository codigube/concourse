// Code generated by mockery v2.2.1. DO NOT EDIT.

package cmocks

import mock "github.com/stretchr/testify/mock"

// NotificationsBus is an autogenerated mock type for the NotificationsBus type
type NotificationsBus struct {
	mock.Mock
}

// Listen provides a mock function with given fields: _a0
func (_m *NotificationsBus) Listen(_a0 string) (chan bool, error) {
	ret := _m.Called(_a0)

	var r0 chan bool
	if rf, ok := ret.Get(0).(func(string) chan bool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlisten provides a mock function with given fields: _a0, _a1
func (_m *NotificationsBus) Unlisten(_a0 string, _a1 chan bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, chan bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
