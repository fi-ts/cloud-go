// Code generated by mockery (devel). DO NOT EDIT.

package project

import (
	project "github.com/fi-ts/cloud-go/api/client/project"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreateProject provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreateProject(params *project.CreateProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...project.ClientOption) (*project.CreateProjectCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *project.CreateProjectCreated
	if rf, ok := ret.Get(0).(func(*project.CreateProjectParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) *project.CreateProjectCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.CreateProjectCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.CreateProjectParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProject provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteProject(params *project.DeleteProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...project.ClientOption) (*project.DeleteProjectOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *project.DeleteProjectOK
	if rf, ok := ret.Get(0).(func(*project.DeleteProjectParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) *project.DeleteProjectOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.DeleteProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.DeleteProjectParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindProject provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindProject(params *project.FindProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...project.ClientOption) (*project.FindProjectOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *project.FindProjectOK
	if rf, ok := ret.Get(0).(func(*project.FindProjectParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) *project.FindProjectOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.FindProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.FindProjectParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListProjects(params *project.ListProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...project.ClientOption) (*project.ListProjectsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *project.ListProjectsOK
	if rf, ok := ret.Get(0).(func(*project.ListProjectsParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) *project.ListProjectsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.ListProjectsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.ListProjectsParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) error); ok {
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

// UpdateProject provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdateProject(params *project.UpdateProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...project.ClientOption) (*project.UpdateProjectOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *project.UpdateProjectOK
	if rf, ok := ret.Get(0).(func(*project.UpdateProjectParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) *project.UpdateProjectOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project.UpdateProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project.UpdateProjectParams, runtime.ClientAuthInfoWriter, ...project.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
