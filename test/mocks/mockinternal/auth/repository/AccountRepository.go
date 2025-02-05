// Code generated by mockery v2.42.1. DO NOT EDIT.

package repository

import (
	context "context"
	aggregate "ddd-sample/internal/auth/aggregate"

	coreaggregate "ddd-sample/internal/core/aggregate"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

type AccountRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountRepository) EXPECT() *AccountRepository_Expecter {
	return &AccountRepository_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, account
func (_m *AccountRepository) Add(ctx context.Context, account *aggregate.Account) error {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregate.Account) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepository_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type AccountRepository_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - account *aggregate.Account
func (_e *AccountRepository_Expecter) Add(ctx interface{}, account interface{}) *AccountRepository_Add_Call {
	return &AccountRepository_Add_Call{Call: _e.mock.On("Add", ctx, account)}
}

func (_c *AccountRepository_Add_Call) Run(run func(ctx context.Context, account *aggregate.Account)) *AccountRepository_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*aggregate.Account))
	})
	return _c
}

func (_c *AccountRepository_Add_Call) Return(_a0 error) *AccountRepository_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_Add_Call) RunAndReturn(run func(context.Context, *aggregate.Account) error) *AccountRepository_Add_Call {
	_c.Call.Return(run)
	return _c
}

// ChangePassword provides a mock function with given fields: ctx, account
func (_m *AccountRepository) ChangePassword(ctx context.Context, account *aggregate.Account) error {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for ChangePassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregate.Account) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepository_ChangePassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangePassword'
type AccountRepository_ChangePassword_Call struct {
	*mock.Call
}

// ChangePassword is a helper method to define mock.On call
//   - ctx context.Context
//   - account *aggregate.Account
func (_e *AccountRepository_Expecter) ChangePassword(ctx interface{}, account interface{}) *AccountRepository_ChangePassword_Call {
	return &AccountRepository_ChangePassword_Call{Call: _e.mock.On("ChangePassword", ctx, account)}
}

func (_c *AccountRepository_ChangePassword_Call) Run(run func(ctx context.Context, account *aggregate.Account)) *AccountRepository_ChangePassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*aggregate.Account))
	})
	return _c
}

func (_c *AccountRepository_ChangePassword_Call) Return(_a0 error) *AccountRepository_ChangePassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_ChangePassword_Call) RunAndReturn(run func(context.Context, *aggregate.Account) error) *AccountRepository_ChangePassword_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: ctx, uid
func (_m *AccountRepository) Find(ctx context.Context, uid string) (*aggregate.Account, error) {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *aggregate.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*aggregate.Account, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *aggregate.Account); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type AccountRepository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - ctx context.Context
//   - uid string
func (_e *AccountRepository_Expecter) Find(ctx interface{}, uid interface{}) *AccountRepository_Find_Call {
	return &AccountRepository_Find_Call{Call: _e.mock.On("Find", ctx, uid)}
}

func (_c *AccountRepository_Find_Call) Run(run func(ctx context.Context, uid string)) *AccountRepository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AccountRepository_Find_Call) Return(_a0 *aggregate.Account, _a1 error) *AccountRepository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepository_Find_Call) RunAndReturn(run func(context.Context, string) (*aggregate.Account, error)) *AccountRepository_Find_Call {
	_c.Call.Return(run)
	return _c
}

// New provides a mock function with given fields: username, password, nowTime
func (_m *AccountRepository) New(username string, password string, nowTime time.Time) *aggregate.Account {
	ret := _m.Called(username, password, nowTime)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 *aggregate.Account
	if rf, ok := ret.Get(0).(func(string, string, time.Time) *aggregate.Account); ok {
		r0 = rf(username, password, nowTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*aggregate.Account)
		}
	}

	return r0
}

// AccountRepository_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type AccountRepository_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - username string
//   - password string
//   - nowTime time.Time
func (_e *AccountRepository_Expecter) New(username interface{}, password interface{}, nowTime interface{}) *AccountRepository_New_Call {
	return &AccountRepository_New_Call{Call: _e.mock.On("New", username, password, nowTime)}
}

func (_c *AccountRepository_New_Call) Run(run func(username string, password string, nowTime time.Time)) *AccountRepository_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(time.Time))
	})
	return _c
}

func (_c *AccountRepository_New_Call) Return(_a0 *aggregate.Account) *AccountRepository_New_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_New_Call) RunAndReturn(run func(string, string, time.Time) *aggregate.Account) *AccountRepository_New_Call {
	_c.Call.Return(run)
	return _c
}

// PubEvent provides a mock function with given fields: _a0
func (_m *AccountRepository) PubEvent(_a0 coreaggregate.CoreAggregate) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for PubEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(coreaggregate.CoreAggregate) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepository_PubEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PubEvent'
type AccountRepository_PubEvent_Call struct {
	*mock.Call
}

// PubEvent is a helper method to define mock.On call
//   - _a0 coreaggregate.CoreAggregate
func (_e *AccountRepository_Expecter) PubEvent(_a0 interface{}) *AccountRepository_PubEvent_Call {
	return &AccountRepository_PubEvent_Call{Call: _e.mock.On("PubEvent", _a0)}
}

func (_c *AccountRepository_PubEvent_Call) Run(run func(_a0 coreaggregate.CoreAggregate)) *AccountRepository_PubEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(coreaggregate.CoreAggregate))
	})
	return _c
}

func (_c *AccountRepository_PubEvent_Call) Return(_a0 error) *AccountRepository_PubEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_PubEvent_Call) RunAndReturn(run func(coreaggregate.CoreAggregate) error) *AccountRepository_PubEvent_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, account
func (_m *AccountRepository) Update(ctx context.Context, account *aggregate.Account) error {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *aggregate.Account) error); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type AccountRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - account *aggregate.Account
func (_e *AccountRepository_Expecter) Update(ctx interface{}, account interface{}) *AccountRepository_Update_Call {
	return &AccountRepository_Update_Call{Call: _e.mock.On("Update", ctx, account)}
}

func (_c *AccountRepository_Update_Call) Run(run func(ctx context.Context, account *aggregate.Account)) *AccountRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*aggregate.Account))
	})
	return _c
}

func (_c *AccountRepository_Update_Call) Return(_a0 error) *AccountRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepository_Update_Call) RunAndReturn(run func(context.Context, *aggregate.Account) error) *AccountRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
