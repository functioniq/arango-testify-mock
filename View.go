// Code generated by mockery v1.0.0. DO NOT EDIT.

package arangomock

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// View is an autogenerated mock type for the View type
type View struct {
	mock.Mock
}

// ArangoSearchView provides a mock function with given fields:
func (_m *View) ArangoSearchView() (driver.ArangoSearchView, error) {
	ret := _m.Called()

	var r0 driver.ArangoSearchView
	if rf, ok := ret.Get(0).(func() driver.ArangoSearchView); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.ArangoSearchView)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Database provides a mock function with given fields:
func (_m *View) Database() driver.Database {
	ret := _m.Called()

	var r0 driver.Database
	if rf, ok := ret.Get(0).(func() driver.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Database)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *View) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Remove provides a mock function with given fields: ctx
func (_m *View) Remove(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *View) Type() driver.ViewType {
	ret := _m.Called()

	var r0 driver.ViewType
	if rf, ok := ret.Get(0).(func() driver.ViewType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(driver.ViewType)
	}

	return r0
}
