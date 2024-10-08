// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"
	errors "hidroponic/internal/errors"
	entities "hidroponic/internal/module/installationconfig/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// FindContainerConfigByID provides a mock function with given fields: ctx, ContainerID
func (_m *MockRepository) FindContainerConfigByID(ctx context.Context, ContainerID string) (*entities.ContainerConfig, *errors.BaseError) {
	ret := _m.Called(ctx, ContainerID)

	if len(ret) == 0 {
		panic("no return value specified for FindContainerConfigByID")
	}

	var r0 *entities.ContainerConfig
	var r1 *errors.BaseError
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.ContainerConfig, *errors.BaseError)); ok {
		return rf(ctx, ContainerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.ContainerConfig); ok {
		r0 = rf(ctx, ContainerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ContainerConfig)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) *errors.BaseError); ok {
		r1 = rf(ctx, ContainerID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.BaseError)
		}
	}

	return r0, r1
}

// MockRepository_FindContainerConfigByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindContainerConfigByID'
type MockRepository_FindContainerConfigByID_Call struct {
	*mock.Call
}

// FindContainerConfigByID is a helper method to define mock.On call
//   - ctx context.Context
//   - ContainerID string
func (_e *MockRepository_Expecter) FindContainerConfigByID(ctx interface{}, ContainerID interface{}) *MockRepository_FindContainerConfigByID_Call {
	return &MockRepository_FindContainerConfigByID_Call{Call: _e.mock.On("FindContainerConfigByID", ctx, ContainerID)}
}

func (_c *MockRepository_FindContainerConfigByID_Call) Run(run func(ctx context.Context, ContainerID string)) *MockRepository_FindContainerConfigByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_FindContainerConfigByID_Call) Return(_a0 *entities.ContainerConfig, _a1 *errors.BaseError) *MockRepository_FindContainerConfigByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindContainerConfigByID_Call) RunAndReturn(run func(context.Context, string) (*entities.ContainerConfig, *errors.BaseError)) *MockRepository_FindContainerConfigByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllContainerConfig provides a mock function with given fields: ctx
func (_m *MockRepository) GetAllContainerConfig(ctx context.Context) ([]entities.ContainerConfig, *errors.BaseError) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllContainerConfig")
	}

	var r0 []entities.ContainerConfig
	var r1 *errors.BaseError
	if rf, ok := ret.Get(0).(func(context.Context) ([]entities.ContainerConfig, *errors.BaseError)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entities.ContainerConfig); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ContainerConfig)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) *errors.BaseError); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.BaseError)
		}
	}

	return r0, r1
}

// MockRepository_GetAllContainerConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllContainerConfig'
type MockRepository_GetAllContainerConfig_Call struct {
	*mock.Call
}

// GetAllContainerConfig is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepository_Expecter) GetAllContainerConfig(ctx interface{}) *MockRepository_GetAllContainerConfig_Call {
	return &MockRepository_GetAllContainerConfig_Call{Call: _e.mock.On("GetAllContainerConfig", ctx)}
}

func (_c *MockRepository_GetAllContainerConfig_Call) Run(run func(ctx context.Context)) *MockRepository_GetAllContainerConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepository_GetAllContainerConfig_Call) Return(_a0 []entities.ContainerConfig, _a1 *errors.BaseError) *MockRepository_GetAllContainerConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetAllContainerConfig_Call) RunAndReturn(run func(context.Context) ([]entities.ContainerConfig, *errors.BaseError)) *MockRepository_GetAllContainerConfig_Call {
	_c.Call.Return(run)
	return _c
}

// GetInstallationConfig provides a mock function with given fields: ctx, ID
func (_m *MockRepository) GetInstallationConfig(ctx context.Context, ID uint) (*entities.InstallationConfig, *errors.BaseError) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetInstallationConfig")
	}

	var r0 *entities.InstallationConfig
	var r1 *errors.BaseError
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*entities.InstallationConfig, *errors.BaseError)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *entities.InstallationConfig); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.InstallationConfig)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) *errors.BaseError); ok {
		r1 = rf(ctx, ID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.BaseError)
		}
	}

	return r0, r1
}

// MockRepository_GetInstallationConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInstallationConfig'
type MockRepository_GetInstallationConfig_Call struct {
	*mock.Call
}

// GetInstallationConfig is a helper method to define mock.On call
//   - ctx context.Context
//   - ID uint
func (_e *MockRepository_Expecter) GetInstallationConfig(ctx interface{}, ID interface{}) *MockRepository_GetInstallationConfig_Call {
	return &MockRepository_GetInstallationConfig_Call{Call: _e.mock.On("GetInstallationConfig", ctx, ID)}
}

func (_c *MockRepository_GetInstallationConfig_Call) Run(run func(ctx context.Context, ID uint)) *MockRepository_GetInstallationConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *MockRepository_GetInstallationConfig_Call) Return(_a0 *entities.InstallationConfig, _a1 *errors.BaseError) *MockRepository_GetInstallationConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetInstallationConfig_Call) RunAndReturn(run func(context.Context, uint) (*entities.InstallationConfig, *errors.BaseError)) *MockRepository_GetInstallationConfig_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateContainerConfig provides a mock function with given fields: ctx, data
func (_m *MockRepository) UpdateContainerConfig(ctx context.Context, data *entities.ContainerConfig) *errors.BaseError {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateContainerConfig")
	}

	var r0 *errors.BaseError
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ContainerConfig) *errors.BaseError); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.BaseError)
		}
	}

	return r0
}

// MockRepository_UpdateContainerConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateContainerConfig'
type MockRepository_UpdateContainerConfig_Call struct {
	*mock.Call
}

// UpdateContainerConfig is a helper method to define mock.On call
//   - ctx context.Context
//   - data *entities.ContainerConfig
func (_e *MockRepository_Expecter) UpdateContainerConfig(ctx interface{}, data interface{}) *MockRepository_UpdateContainerConfig_Call {
	return &MockRepository_UpdateContainerConfig_Call{Call: _e.mock.On("UpdateContainerConfig", ctx, data)}
}

func (_c *MockRepository_UpdateContainerConfig_Call) Run(run func(ctx context.Context, data *entities.ContainerConfig)) *MockRepository_UpdateContainerConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.ContainerConfig))
	})
	return _c
}

func (_c *MockRepository_UpdateContainerConfig_Call) Return(_a0 *errors.BaseError) *MockRepository_UpdateContainerConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_UpdateContainerConfig_Call) RunAndReturn(run func(context.Context, *entities.ContainerConfig) *errors.BaseError) *MockRepository_UpdateContainerConfig_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateInstallationConfig provides a mock function with given fields: ctx, data
func (_m *MockRepository) UpdateInstallationConfig(ctx context.Context, data *entities.InstallationConfig) *errors.BaseError {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateInstallationConfig")
	}

	var r0 *errors.BaseError
	if rf, ok := ret.Get(0).(func(context.Context, *entities.InstallationConfig) *errors.BaseError); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.BaseError)
		}
	}

	return r0
}

// MockRepository_UpdateInstallationConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateInstallationConfig'
type MockRepository_UpdateInstallationConfig_Call struct {
	*mock.Call
}

// UpdateInstallationConfig is a helper method to define mock.On call
//   - ctx context.Context
//   - data *entities.InstallationConfig
func (_e *MockRepository_Expecter) UpdateInstallationConfig(ctx interface{}, data interface{}) *MockRepository_UpdateInstallationConfig_Call {
	return &MockRepository_UpdateInstallationConfig_Call{Call: _e.mock.On("UpdateInstallationConfig", ctx, data)}
}

func (_c *MockRepository_UpdateInstallationConfig_Call) Run(run func(ctx context.Context, data *entities.InstallationConfig)) *MockRepository_UpdateInstallationConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.InstallationConfig))
	})
	return _c
}

func (_c *MockRepository_UpdateInstallationConfig_Call) Return(_a0 *errors.BaseError) *MockRepository_UpdateInstallationConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_UpdateInstallationConfig_Call) RunAndReturn(run func(context.Context, *entities.InstallationConfig) *errors.BaseError) *MockRepository_UpdateInstallationConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
