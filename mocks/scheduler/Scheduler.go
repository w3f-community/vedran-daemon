// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	gocron "github.com/go-co-op/gocron"
	mock "github.com/stretchr/testify/mock"
)

// Scheduler is an autogenerated mock type for the Scheduler type
type Scheduler struct {
	mock.Mock
}

// Do provides a mock function with given fields: jobFun, params
func (_m *Scheduler) Do(jobFun interface{}, params ...interface{}) (*gocron.Job, error) {
	var _ca []interface{}
	_ca = append(_ca, jobFun)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 *gocron.Job
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gocron.Job); ok {
		r0 = rf(jobFun, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gocron.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(jobFun, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Every provides a mock function with given fields: interval
func (_m *Scheduler) Every(interval uint64) *gocron.Scheduler {
	ret := _m.Called(interval)

	var r0 *gocron.Scheduler
	if rf, ok := ret.Get(0).(func(uint64) *gocron.Scheduler); ok {
		r0 = rf(interval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gocron.Scheduler)
		}
	}

	return r0
}

// StartBlocking provides a mock function with given fields:
func (_m *Scheduler) StartBlocking() {
	_m.Called()
}
