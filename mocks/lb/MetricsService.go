// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	node "github.com/NodeFactoryIo/vedran-daemon/internal/node"
)

// MetricsService is an autogenerated mock type for the MetricsService type
type MetricsService struct {
	mock.Mock
}

// Send provides a mock function with given fields: client
func (_m *MetricsService) Send(client node.Client) (*http.Response, error) {
	ret := _m.Called(client)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(node.Client) *http.Response); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(node.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
