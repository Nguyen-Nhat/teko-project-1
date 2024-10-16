// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"
	database "library/internal/database"

	mock "github.com/stretchr/testify/mock"

	req "library/internal/dto/req"
)

// IGenreService is an autogenerated mock type for the IGenreService type
type IGenreService struct {
	mock.Mock
}

// CreateGenre provides a mock function with given fields: ctx, data
func (_m *IGenreService) CreateGenre(ctx context.Context, data *req.GenrePostDto) (*database.Genre, int, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for CreateGenre")
	}

	var r0 *database.Genre
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *req.GenrePostDto) (*database.Genre, int, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *req.GenrePostDto) *database.Genre); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Genre)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *req.GenrePostDto) int); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *req.GenrePostDto) error); ok {
		r2 = rf(ctx, data)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetGenreById provides a mock function with given fields: ctx, id
func (_m *IGenreService) GetGenreById(ctx context.Context, id int) (*database.Genre, int, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetGenreById")
	}

	var r0 *database.Genre
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*database.Genre, int, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *database.Genre); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Genre)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) int); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewIGenreService creates a new instance of IGenreService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIGenreService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IGenreService {
	mock := &IGenreService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
