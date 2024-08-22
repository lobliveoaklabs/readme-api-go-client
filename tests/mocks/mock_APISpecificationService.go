// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	readme "github.com/liveoaklabs/readme-api-go-client/readme"
	mock "github.com/stretchr/testify/mock"
)

// MockAPISpecificationService is an autogenerated mock type for the APISpecificationService type
type MockAPISpecificationService struct {
	mock.Mock
}

type MockAPISpecificationService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAPISpecificationService) EXPECT() *MockAPISpecificationService_Expecter {
	return &MockAPISpecificationService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: definition, options
func (_m *MockAPISpecificationService) Create(definition string, options ...readme.RequestOptions) (readme.APISpecificationSaved, *readme.APIResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, definition)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 readme.APISpecificationSaved
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) (readme.APISpecificationSaved, *readme.APIResponse, error)); ok {
		return rf(definition, options...)
	}
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) readme.APISpecificationSaved); ok {
		r0 = rf(definition, options...)
	} else {
		r0 = ret.Get(0).(readme.APISpecificationSaved)
	}

	if rf, ok := ret.Get(1).(func(string, ...readme.RequestOptions) *readme.APIResponse); ok {
		r1 = rf(definition, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, ...readme.RequestOptions) error); ok {
		r2 = rf(definition, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAPISpecificationService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAPISpecificationService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - definition string
//   - options ...readme.RequestOptions
func (_e *MockAPISpecificationService_Expecter) Create(definition interface{}, options ...interface{}) *MockAPISpecificationService_Create_Call {
	return &MockAPISpecificationService_Create_Call{Call: _e.mock.On("Create",
		append([]interface{}{definition}, options...)...)}
}

func (_c *MockAPISpecificationService_Create_Call) Run(run func(definition string, options ...readme.RequestOptions)) *MockAPISpecificationService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAPISpecificationService_Create_Call) Return(_a0 readme.APISpecificationSaved, _a1 *readme.APIResponse, _a2 error) *MockAPISpecificationService_Create_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAPISpecificationService_Create_Call) RunAndReturn(run func(string, ...readme.RequestOptions) (readme.APISpecificationSaved, *readme.APIResponse, error)) *MockAPISpecificationService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: specID
func (_m *MockAPISpecificationService) Delete(specID string) (bool, *readme.APIResponse, error) {
	ret := _m.Called(specID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 bool
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (bool, *readme.APIResponse, error)); ok {
		return rf(specID)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(specID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) *readme.APIResponse); ok {
		r1 = rf(specID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(specID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAPISpecificationService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAPISpecificationService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - specID string
func (_e *MockAPISpecificationService_Expecter) Delete(specID interface{}) *MockAPISpecificationService_Delete_Call {
	return &MockAPISpecificationService_Delete_Call{Call: _e.mock.On("Delete", specID)}
}

func (_c *MockAPISpecificationService_Delete_Call) Run(run func(specID string)) *MockAPISpecificationService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAPISpecificationService_Delete_Call) Return(_a0 bool, _a1 *readme.APIResponse, _a2 error) *MockAPISpecificationService_Delete_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAPISpecificationService_Delete_Call) RunAndReturn(run func(string) (bool, *readme.APIResponse, error)) *MockAPISpecificationService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: specID, options
func (_m *MockAPISpecificationService) Get(specID string, options ...readme.RequestOptions) (readme.APISpecification, *readme.APIResponse, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, specID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 readme.APISpecification
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) (readme.APISpecification, *readme.APIResponse, error)); ok {
		return rf(specID, options...)
	}
	if rf, ok := ret.Get(0).(func(string, ...readme.RequestOptions) readme.APISpecification); ok {
		r0 = rf(specID, options...)
	} else {
		r0 = ret.Get(0).(readme.APISpecification)
	}

	if rf, ok := ret.Get(1).(func(string, ...readme.RequestOptions) *readme.APIResponse); ok {
		r1 = rf(specID, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, ...readme.RequestOptions) error); ok {
		r2 = rf(specID, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAPISpecificationService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAPISpecificationService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - specID string
//   - options ...readme.RequestOptions
func (_e *MockAPISpecificationService_Expecter) Get(specID interface{}, options ...interface{}) *MockAPISpecificationService_Get_Call {
	return &MockAPISpecificationService_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{specID}, options...)...)}
}

func (_c *MockAPISpecificationService_Get_Call) Run(run func(specID string, options ...readme.RequestOptions)) *MockAPISpecificationService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockAPISpecificationService_Get_Call) Return(_a0 readme.APISpecification, _a1 *readme.APIResponse, _a2 error) *MockAPISpecificationService_Get_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAPISpecificationService_Get_Call) RunAndReturn(run func(string, ...readme.RequestOptions) (readme.APISpecification, *readme.APIResponse, error)) *MockAPISpecificationService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: _a0
func (_m *MockAPISpecificationService) GetAll(_a0 ...readme.RequestOptions) ([]readme.APISpecification, *readme.APIResponse, error) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []readme.APISpecification
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(...readme.RequestOptions) ([]readme.APISpecification, *readme.APIResponse, error)); ok {
		return rf(_a0...)
	}
	if rf, ok := ret.Get(0).(func(...readme.RequestOptions) []readme.APISpecification); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]readme.APISpecification)
		}
	}

	if rf, ok := ret.Get(1).(func(...readme.RequestOptions) *readme.APIResponse); ok {
		r1 = rf(_a0...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(...readme.RequestOptions) error); ok {
		r2 = rf(_a0...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAPISpecificationService_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockAPISpecificationService_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - _a0 ...readme.RequestOptions
func (_e *MockAPISpecificationService_Expecter) GetAll(_a0 ...interface{}) *MockAPISpecificationService_GetAll_Call {
	return &MockAPISpecificationService_GetAll_Call{Call: _e.mock.On("GetAll",
		append([]interface{}{}, _a0...)...)}
}

func (_c *MockAPISpecificationService_GetAll_Call) Run(run func(_a0 ...readme.RequestOptions)) *MockAPISpecificationService_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]readme.RequestOptions, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(readme.RequestOptions)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockAPISpecificationService_GetAll_Call) Return(_a0 []readme.APISpecification, _a1 *readme.APIResponse, _a2 error) *MockAPISpecificationService_GetAll_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAPISpecificationService_GetAll_Call) RunAndReturn(run func(...readme.RequestOptions) ([]readme.APISpecification, *readme.APIResponse, error)) *MockAPISpecificationService_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: specID, definition
func (_m *MockAPISpecificationService) Update(specID string, definition string) (readme.APISpecificationSaved, *readme.APIResponse, error) {
	ret := _m.Called(specID, definition)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 readme.APISpecificationSaved
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (readme.APISpecificationSaved, *readme.APIResponse, error)); ok {
		return rf(specID, definition)
	}
	if rf, ok := ret.Get(0).(func(string, string) readme.APISpecificationSaved); ok {
		r0 = rf(specID, definition)
	} else {
		r0 = ret.Get(0).(readme.APISpecificationSaved)
	}

	if rf, ok := ret.Get(1).(func(string, string) *readme.APIResponse); ok {
		r1 = rf(specID, definition)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(specID, definition)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAPISpecificationService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAPISpecificationService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - specID string
//   - definition string
func (_e *MockAPISpecificationService_Expecter) Update(specID interface{}, definition interface{}) *MockAPISpecificationService_Update_Call {
	return &MockAPISpecificationService_Update_Call{Call: _e.mock.On("Update", specID, definition)}
}

func (_c *MockAPISpecificationService_Update_Call) Run(run func(specID string, definition string)) *MockAPISpecificationService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockAPISpecificationService_Update_Call) Return(_a0 readme.APISpecificationSaved, _a1 *readme.APIResponse, _a2 error) *MockAPISpecificationService_Update_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAPISpecificationService_Update_Call) RunAndReturn(run func(string, string) (readme.APISpecificationSaved, *readme.APIResponse, error)) *MockAPISpecificationService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UploadDefinition provides a mock function with given fields: method, content, url, version, response
func (_m *MockAPISpecificationService) UploadDefinition(method string, content string, url string, version string, response interface{}) (interface{}, *readme.APIResponse, error) {
	ret := _m.Called(method, content, url, version, response)

	if len(ret) == 0 {
		panic("no return value specified for UploadDefinition")
	}

	var r0 interface{}
	var r1 *readme.APIResponse
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, interface{}) (interface{}, *readme.APIResponse, error)); ok {
		return rf(method, content, url, version, response)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string, interface{}) interface{}); ok {
		r0 = rf(method, content, url, version, response)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string, interface{}) *readme.APIResponse); ok {
		r1 = rf(method, content, url, version, response)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*readme.APIResponse)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string, string, string, interface{}) error); ok {
		r2 = rf(method, content, url, version, response)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockAPISpecificationService_UploadDefinition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadDefinition'
type MockAPISpecificationService_UploadDefinition_Call struct {
	*mock.Call
}

// UploadDefinition is a helper method to define mock.On call
//   - method string
//   - content string
//   - url string
//   - version string
//   - response interface{}
func (_e *MockAPISpecificationService_Expecter) UploadDefinition(method interface{}, content interface{}, url interface{}, version interface{}, response interface{}) *MockAPISpecificationService_UploadDefinition_Call {
	return &MockAPISpecificationService_UploadDefinition_Call{Call: _e.mock.On("UploadDefinition", method, content, url, version, response)}
}

func (_c *MockAPISpecificationService_UploadDefinition_Call) Run(run func(method string, content string, url string, version string, response interface{})) *MockAPISpecificationService_UploadDefinition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(interface{}))
	})
	return _c
}

func (_c *MockAPISpecificationService_UploadDefinition_Call) Return(_a0 interface{}, _a1 *readme.APIResponse, _a2 error) *MockAPISpecificationService_UploadDefinition_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockAPISpecificationService_UploadDefinition_Call) RunAndReturn(run func(string, string, string, string, interface{}) (interface{}, *readme.APIResponse, error)) *MockAPISpecificationService_UploadDefinition_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAPISpecificationService creates a new instance of MockAPISpecificationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAPISpecificationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAPISpecificationService {
	mock := &MockAPISpecificationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
