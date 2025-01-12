// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"
	model "student/internal/model"

	mock "github.com/stretchr/testify/mock"

	req "student/internal/dto/req"

	res "student/internal/dto/res"
)

// IStudentRepository is an autogenerated mock type for the IStudentRepository type
type IStudentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, student
func (_m *IStudentRepository) Create(ctx context.Context, student *model.Student) error {
	ret := _m.Called(ctx, student)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Student) error); ok {
		r0 = rf(ctx, student)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *IStudentRepository) FindByID(ctx context.Context, id int) (*model.Student, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *model.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*model.Student, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *model.Student); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Student)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPageByUniIdAndEnrollYear provides a mock function with given fields: ctx, universityId, enrollYear, page
func (_m *IStudentRepository) FindPageByUniIdAndEnrollYear(ctx context.Context, universityId int, enrollYear int, page req.PageInfo) (*res.PageResult[model.Student], error) {
	ret := _m.Called(ctx, universityId, enrollYear, page)

	if len(ret) == 0 {
		panic("no return value specified for FindPageByUniIdAndEnrollYear")
	}

	var r0 *res.PageResult[model.Student]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, req.PageInfo) (*res.PageResult[model.Student], error)); ok {
		return rf(ctx, universityId, enrollYear, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int, req.PageInfo) *res.PageResult[model.Student]); ok {
		r0 = rf(ctx, universityId, enrollYear, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*res.PageResult[model.Student])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int, req.PageInfo) error); ok {
		r1 = rf(ctx, universityId, enrollYear, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIStudentRepository creates a new instance of IStudentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIStudentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IStudentRepository {
	mock := &IStudentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
