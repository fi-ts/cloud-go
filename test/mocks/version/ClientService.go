// Code generated by mockery v2.8.0. DO NOT EDIT.

package version

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	version "github.com/fi-ts/cloud-go/api/client/version"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// Info provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Info(params *version.InfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...version.ClientOption) (*version.InfoOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *version.InfoOK
	if rf, ok := ret.Get(0).(func(*version.InfoParams, runtime.ClientAuthInfoWriter, ...version.ClientOption) *version.InfoOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.InfoOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*version.InfoParams, runtime.ClientAuthInfoWriter, ...version.ClientOption) error); ok {
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
