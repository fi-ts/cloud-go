// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package database

import (
	database "github.com/fi-ts/cloud-go/api/client/database"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreatePostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreatePostgres(params *database.CreatePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.CreatePostgresCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.CreatePostgresCreated
	if rf, ok := ret.Get(0).(func(*database.CreatePostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.CreatePostgresCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.CreatePostgresCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.CreatePostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePostgresBackupConfig provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) CreatePostgresBackupConfig(params *database.CreatePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.CreatePostgresBackupConfigCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.CreatePostgresBackupConfigCreated
	if rf, ok := ret.Get(0).(func(*database.CreatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.CreatePostgresBackupConfigCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.CreatePostgresBackupConfigCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.CreatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeletePostgres(params *database.DeletePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.DeletePostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.DeletePostgresOK
	if rf, ok := ret.Get(0).(func(*database.DeletePostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.DeletePostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.DeletePostgresOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.DeletePostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePostgresBackupConfig provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeletePostgresBackupConfig(params *database.DeletePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.DeletePostgresBackupConfigOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.DeletePostgresBackupConfigOK
	if rf, ok := ret.Get(0).(func(*database.DeletePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.DeletePostgresBackupConfigOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.DeletePostgresBackupConfigOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.DeletePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) FindPostgres(params *database.FindPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.FindPostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.FindPostgresOK
	if rf, ok := ret.Get(0).(func(*database.FindPostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.FindPostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.FindPostgresOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.FindPostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupConfig provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetBackupConfig(params *database.GetBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.GetBackupConfigOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.GetBackupConfigOK
	if rf, ok := ret.Get(0).(func(*database.GetBackupConfigParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.GetBackupConfigOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.GetBackupConfigOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.GetBackupConfigParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgres(params *database.GetPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.GetPostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.GetPostgresOK
	if rf, ok := ret.Get(0).(func(*database.GetPostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.GetPostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.GetPostgresOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.GetPostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgresBackups provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgresBackups(params *database.GetPostgresBackupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.GetPostgresBackupsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.GetPostgresBackupsOK
	if rf, ok := ret.Get(0).(func(*database.GetPostgresBackupsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.GetPostgresBackupsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.GetPostgresBackupsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.GetPostgresBackupsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgresPartitions provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgresPartitions(params *database.GetPostgresPartitionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.GetPostgresPartitionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.GetPostgresPartitionsOK
	if rf, ok := ret.Get(0).(func(*database.GetPostgresPartitionsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.GetPostgresPartitionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.GetPostgresPartitionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.GetPostgresPartitionsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgresSecrets provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgresSecrets(params *database.GetPostgresSecretsParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.GetPostgresSecretsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.GetPostgresSecretsOK
	if rf, ok := ret.Get(0).(func(*database.GetPostgresSecretsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.GetPostgresSecretsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.GetPostgresSecretsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.GetPostgresSecretsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostgresVersions provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPostgresVersions(params *database.GetPostgresVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.GetPostgresVersionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.GetPostgresVersionsOK
	if rf, ok := ret.Get(0).(func(*database.GetPostgresVersionsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.GetPostgresVersionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.GetPostgresVersionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.GetPostgresVersionsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListPostgres(params *database.ListPostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.ListPostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.ListPostgresOK
	if rf, ok := ret.Get(0).(func(*database.ListPostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.ListPostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.ListPostgresOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.ListPostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPostgresBackupConfigs provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) ListPostgresBackupConfigs(params *database.ListPostgresBackupConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.ListPostgresBackupConfigsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.ListPostgresBackupConfigsOK
	if rf, ok := ret.Get(0).(func(*database.ListPostgresBackupConfigsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.ListPostgresBackupConfigsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.ListPostgresBackupConfigsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.ListPostgresBackupConfigsParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
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

// UpdatePostgres provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdatePostgres(params *database.UpdatePostgresParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.UpdatePostgresOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.UpdatePostgresOK
	if rf, ok := ret.Get(0).(func(*database.UpdatePostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.UpdatePostgresOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.UpdatePostgresOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.UpdatePostgresParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePostgresBackupConfig provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) UpdatePostgresBackupConfig(params *database.UpdatePostgresBackupConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...database.ClientOption) (*database.UpdatePostgresBackupConfigOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *database.UpdatePostgresBackupConfigOK
	if rf, ok := ret.Get(0).(func(*database.UpdatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) *database.UpdatePostgresBackupConfigOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.UpdatePostgresBackupConfigOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*database.UpdatePostgresBackupConfigParams, runtime.ClientAuthInfoWriter, ...database.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
