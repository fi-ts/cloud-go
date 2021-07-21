// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package health

import (
	health "github.com/fi-ts/cloud-go/api/client/health"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// Health provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Health(params *health.HealthParams, authInfo runtime.ClientAuthInfoWriter, opts ...health.ClientOption) (*health.HealthOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *health.HealthOK
	if rf, ok := ret.Get(0).(func(*health.HealthParams, runtime.ClientAuthInfoWriter, ...health.ClientOption) *health.HealthOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*health.HealthOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*health.HealthParams, runtime.ClientAuthInfoWriter, ...health.ClientOption) error); ok {
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
