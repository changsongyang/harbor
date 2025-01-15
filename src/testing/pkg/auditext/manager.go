// Code generated by mockery v2.51.0. DO NOT EDIT.

package auditext

import (
	context "context"

	model "github.com/goharbor/harbor/src/pkg/auditext/model"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *Manager) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) (int64, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, audit
func (_m *Manager) Create(ctx context.Context, audit *model.AuditLogExt) (int64, error) {
	ret := _m.Called(ctx, audit)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.AuditLogExt) (int64, error)); ok {
		return rf(ctx, audit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.AuditLogExt) int64); ok {
		r0 = rf(ctx, audit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.AuditLogExt) error); ok {
		r1 = rf(ctx, audit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Manager) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *Manager) Get(ctx context.Context, id int64) (*model.AuditLogExt, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.AuditLogExt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*model.AuditLogExt, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.AuditLogExt); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AuditLogExt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *Manager) List(ctx context.Context, query *q.Query) ([]*model.AuditLogExt, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*model.AuditLogExt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*model.AuditLogExt, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*model.AuditLogExt); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.AuditLogExt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Purge provides a mock function with given fields: ctx, retentionHour, includeOperations, dryRun
func (_m *Manager) Purge(ctx context.Context, retentionHour int, includeOperations []string, dryRun bool) (int64, error) {
	ret := _m.Called(ctx, retentionHour, includeOperations, dryRun)

	if len(ret) == 0 {
		panic("no return value specified for Purge")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []string, bool) (int64, error)); ok {
		return rf(ctx, retentionHour, includeOperations, dryRun)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, []string, bool) int64); ok {
		r0 = rf(ctx, retentionHour, includeOperations, dryRun)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, []string, bool) error); ok {
		r1 = rf(ctx, retentionHour, includeOperations, dryRun)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUsername provides a mock function with given fields: ctx, username, replaceWith
func (_m *Manager) UpdateUsername(ctx context.Context, username string, replaceWith string) error {
	ret := _m.Called(ctx, username, replaceWith)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUsername")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, username, replaceWith)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}