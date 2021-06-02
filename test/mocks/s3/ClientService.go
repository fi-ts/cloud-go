// Code generated by mockery v2.7.4. DO NOT EDIT.

package s3

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	s3 "github.com/fi-ts/cloud-go/api/client/s3"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// Creates3 provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Creates3(params *s3.Creates3Params, authInfo runtime.ClientAuthInfoWriter, opts ...s3.ClientOption) (*s3.Creates3OK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *s3.Creates3OK
	if rf, ok := ret.Get(0).(func(*s3.Creates3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) *s3.Creates3OK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.Creates3OK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.Creates3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deletes3 provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Deletes3(params *s3.Deletes3Params, authInfo runtime.ClientAuthInfoWriter, opts ...s3.ClientOption) (*s3.Deletes3OK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *s3.Deletes3OK
	if rf, ok := ret.Get(0).(func(*s3.Deletes3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) *s3.Deletes3OK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.Deletes3OK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.Deletes3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Gets3 provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Gets3(params *s3.Gets3Params, authInfo runtime.ClientAuthInfoWriter, opts ...s3.ClientOption) (*s3.Gets3OK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *s3.Gets3OK
	if rf, ok := ret.Get(0).(func(*s3.Gets3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) *s3.Gets3OK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.Gets3OK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.Gets3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lists3 provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Lists3(params *s3.Lists3Params, authInfo runtime.ClientAuthInfoWriter, opts ...s3.ClientOption) (*s3.Lists3OK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *s3.Lists3OK
	if rf, ok := ret.Get(0).(func(*s3.Lists3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) *s3.Lists3OK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.Lists3OK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.Lists3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lists3partitions provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Lists3partitions(params *s3.Lists3partitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...s3.ClientOption) (*s3.Lists3partitionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *s3.Lists3partitionsOK
	if rf, ok := ret.Get(0).(func(*s3.Lists3partitionsParams, runtime.ClientAuthInfoWriter, ...s3.ClientOption) *s3.Lists3partitionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.Lists3partitionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.Lists3partitionsParams, runtime.ClientAuthInfoWriter, ...s3.ClientOption) error); ok {
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

// Updates3 provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) Updates3(params *s3.Updates3Params, authInfo runtime.ClientAuthInfoWriter, opts ...s3.ClientOption) (*s3.Updates3OK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *s3.Updates3OK
	if rf, ok := ret.Get(0).(func(*s3.Updates3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) *s3.Updates3OK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.Updates3OK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*s3.Updates3Params, runtime.ClientAuthInfoWriter, ...s3.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
