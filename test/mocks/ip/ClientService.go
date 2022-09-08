// Code generated by mockery v2.14.0. DO NOT EDIT.

package ip

import (
	clientip "github.com/fi-ts/cloud-go/api/client/ip"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// AllocateIP provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) AllocateIP(params *clientip.AllocateIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientip.ClientOption) (*clientip.AllocateIPCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientip.AllocateIPCreated
	if rf, ok := ret.Get(0).(func(*clientip.AllocateIPParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) *clientip.AllocateIPCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientip.AllocateIPCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientip.AllocateIPParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIPs provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindIPs(params *clientip.FindIPsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientip.ClientOption) (*clientip.FindIPsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientip.FindIPsOK
	if rf, ok := ret.Get(0).(func(*clientip.FindIPsParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) *clientip.FindIPsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientip.FindIPsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientip.FindIPsParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FreeIP provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FreeIP(params *clientip.FreeIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientip.ClientOption) (*clientip.FreeIPOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientip.FreeIPOK
	if rf, ok := ret.Get(0).(func(*clientip.FreeIPParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) *clientip.FreeIPOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientip.FreeIPOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientip.FreeIPParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIP provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetIP(params *clientip.GetIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientip.ClientOption) (*clientip.GetIPOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientip.GetIPOK
	if rf, ok := ret.Get(0).(func(*clientip.GetIPParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) *clientip.GetIPOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientip.GetIPOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientip.GetIPParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIPs provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListIPs(params *clientip.ListIPsParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientip.ClientOption) (*clientip.ListIPsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientip.ListIPsOK
	if rf, ok := ret.Get(0).(func(*clientip.ListIPsParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) *clientip.ListIPsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientip.ListIPsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientip.ListIPsParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) error); ok {
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

// UpdateIP provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateIP(params *clientip.UpdateIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...clientip.ClientOption) (*clientip.UpdateIPOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *clientip.UpdateIPOK
	if rf, ok := ret.Get(0).(func(*clientip.UpdateIPParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) *clientip.UpdateIPOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientip.UpdateIPOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*clientip.UpdateIPParams, runtime.ClientAuthInfoWriter, ...clientip.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientService(t mockConstructorTestingTNewClientService) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
