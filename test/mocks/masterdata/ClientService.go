// Code generated by mockery v2.30.1. DO NOT EDIT.

package masterdata

import (
	clientmasterdata "github.com/fi-ts/cloud-go/api/client/masterdata"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetMasterdata provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetMasterdata(params *clientmasterdata.GetMasterdataParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientmasterdata.ClientOption) (*clientmasterdata.GetMasterdataOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientmasterdata.GetMasterdataOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*clientmasterdata.GetMasterdataParams, runtime.ClientAuthInfoWriter, ...clientmasterdata.ClientOption) (*clientmasterdata.GetMasterdataOK, error)); ok {
		return rf(params, authInfo, opts...)
	}
	if rf, ok := ret.Get(0).(func(*clientmasterdata.GetMasterdataParams, runtime.ClientAuthInfoWriter, ...clientmasterdata.ClientOption) *clientmasterdata.GetMasterdataOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientmasterdata.GetMasterdataOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*clientmasterdata.GetMasterdataParams, runtime.ClientAuthInfoWriter, ...clientmasterdata.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
