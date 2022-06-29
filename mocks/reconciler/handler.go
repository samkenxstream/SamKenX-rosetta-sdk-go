// Code generated by mockery v2.13.1. DO NOT EDIT.

package reconciler

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// ReconciliationExempt provides a mock function with given fields: ctx, reconciliationType, account, currency, computedBalance, liveBalance, block, exemption
func (_m *Handler) ReconciliationExempt(ctx context.Context, reconciliationType string, account *types.AccountIdentifier, currency *types.Currency, computedBalance string, liveBalance string, block *types.BlockIdentifier, exemption *types.BalanceExemption) error {
	ret := _m.Called(ctx, reconciliationType, account, currency, computedBalance, liveBalance, block, exemption)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *types.AccountIdentifier, *types.Currency, string, string, *types.BlockIdentifier, *types.BalanceExemption) error); ok {
		r0 = rf(ctx, reconciliationType, account, currency, computedBalance, liveBalance, block, exemption)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReconciliationFailed provides a mock function with given fields: ctx, reconciliationType, account, currency, computedBalance, liveBalance, block
func (_m *Handler) ReconciliationFailed(ctx context.Context, reconciliationType string, account *types.AccountIdentifier, currency *types.Currency, computedBalance string, liveBalance string, block *types.BlockIdentifier) error {
	ret := _m.Called(ctx, reconciliationType, account, currency, computedBalance, liveBalance, block)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *types.AccountIdentifier, *types.Currency, string, string, *types.BlockIdentifier) error); ok {
		r0 = rf(ctx, reconciliationType, account, currency, computedBalance, liveBalance, block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReconciliationSkipped provides a mock function with given fields: ctx, reconciliationType, account, currency, cause
func (_m *Handler) ReconciliationSkipped(ctx context.Context, reconciliationType string, account *types.AccountIdentifier, currency *types.Currency, cause string) error {
	ret := _m.Called(ctx, reconciliationType, account, currency, cause)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *types.AccountIdentifier, *types.Currency, string) error); ok {
		r0 = rf(ctx, reconciliationType, account, currency, cause)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReconciliationSucceeded provides a mock function with given fields: ctx, reconciliationType, account, currency, balance, block
func (_m *Handler) ReconciliationSucceeded(ctx context.Context, reconciliationType string, account *types.AccountIdentifier, currency *types.Currency, balance string, block *types.BlockIdentifier) error {
	ret := _m.Called(ctx, reconciliationType, account, currency, balance, block)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *types.AccountIdentifier, *types.Currency, string, *types.BlockIdentifier) error); ok {
		r0 = rf(ctx, reconciliationType, account, currency, balance, block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t mockConstructorTestingTNewHandler) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
