// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package tenant

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	tenant "github.com/fi-ts/cloud-go/api/client/tenant"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetTenant provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetTenant(params *tenant.GetTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...tenant.ClientOption) (*tenant.GetTenantOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *tenant.GetTenantOK
	if rf, ok := ret.Get(0).(func(*tenant.GetTenantParams, runtime.ClientAuthInfoWriter, ...tenant.ClientOption) *tenant.GetTenantOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tenant.GetTenantOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*tenant.GetTenantParams, runtime.ClientAuthInfoWriter, ...tenant.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTenants provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListTenants(params *tenant.ListTenantsParams, authInfo runtime.ClientAuthInfoWriter, opts ...tenant.ClientOption) (*tenant.ListTenantsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *tenant.ListTenantsOK
	if rf, ok := ret.Get(0).(func(*tenant.ListTenantsParams, runtime.ClientAuthInfoWriter, ...tenant.ClientOption) *tenant.ListTenantsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tenant.ListTenantsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*tenant.ListTenantsParams, runtime.ClientAuthInfoWriter, ...tenant.ClientOption) error); ok {
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

// UpdateTenant provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateTenant(params *tenant.UpdateTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...tenant.ClientOption) (*tenant.UpdateTenantOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *tenant.UpdateTenantOK
	if rf, ok := ret.Get(0).(func(*tenant.UpdateTenantParams, runtime.ClientAuthInfoWriter, ...tenant.ClientOption) *tenant.UpdateTenantOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tenant.UpdateTenantOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*tenant.UpdateTenantParams, runtime.ClientAuthInfoWriter, ...tenant.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
