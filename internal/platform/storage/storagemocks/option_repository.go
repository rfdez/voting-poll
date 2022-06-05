// Code generated by mockery v2.10.4. DO NOT EDIT.

package storagemocks

import (
	context "context"

	voting "github.com/rfdez/voting-poll/internal"
	mock "github.com/stretchr/testify/mock"
)

// OptionRepository is an autogenerated mock type for the OptionRepository type
type OptionRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: _a0, _a1
func (_m *OptionRepository) Find(_a0 context.Context, _a1 voting.OptionID) (voting.Option, error) {
	ret := _m.Called(_a0, _a1)

	var r0 voting.Option
	if rf, ok := ret.Get(0).(func(context.Context, voting.OptionID) voting.Option); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(voting.Option)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, voting.OptionID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0, _a1
func (_m *OptionRepository) Save(_a0 context.Context, _a1 voting.Option) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, voting.Option) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
