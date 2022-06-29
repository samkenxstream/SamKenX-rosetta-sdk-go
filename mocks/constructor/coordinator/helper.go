// Code generated by mockery v2.13.1. DO NOT EDIT.

package coordinator

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	keys "github.com/coinbase/rosetta-sdk-go/keys"
	database "github.com/coinbase/rosetta-sdk-go/storage/database"
	types "github.com/coinbase/rosetta-sdk-go/types"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// AllAccounts provides a mock function with given fields: _a0, _a1
func (_m *Helper) AllAccounts(_a0 context.Context, _a1 database.Transaction) ([]*types.AccountIdentifier, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.AccountIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction) []*types.AccountIdentifier); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.AccountIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Balance provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Balance(_a0 context.Context, _a1 database.Transaction, _a2 *types.AccountIdentifier, _a3 *types.Currency) (*types.Amount, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Amount
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency) *types.Amount); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Amount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Broadcast provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8
func (_m *Helper) Broadcast(_a0 context.Context, _a1 database.Transaction, _a2 string, _a3 *types.NetworkIdentifier, _a4 []*types.Operation, _a5 *types.TransactionIdentifier, _a6 string, _a7 int64, _a8 map[string]interface{}) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, string, *types.NetworkIdentifier, []*types.Operation, *types.TransactionIdentifier, string, int64, map[string]interface{}) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BroadcastAll provides a mock function with given fields: _a0
func (_m *Helper) BroadcastAll(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Coins provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Coins(_a0 context.Context, _a1 database.Transaction, _a2 *types.AccountIdentifier, _a3 *types.Currency) ([]*types.Coin, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*types.Coin
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency) []*types.Coin); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Coin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, *types.AccountIdentifier, *types.Currency) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Combine provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Combine(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 string, _a3 []*types.Signature) (string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, string, []*types.Signature) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, string, []*types.Signature) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DatabaseTransaction provides a mock function with given fields: _a0
func (_m *Helper) DatabaseTransaction(_a0 context.Context) database.Transaction {
	ret := _m.Called(_a0)

	var r0 database.Transaction
	if rf, ok := ret.Get(0).(func(context.Context) database.Transaction); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Transaction)
		}
	}

	return r0
}

// Derive provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Derive(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 *types.PublicKey, _a3 map[string]interface{}) (*types.AccountIdentifier, map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.AccountIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) *types.AccountIdentifier); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.AccountIdentifier)
		}
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, *types.PublicKey, map[string]interface{}) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetBlob provides a mock function with given fields: ctx, dbTx, key
func (_m *Helper) GetBlob(ctx context.Context, dbTx database.Transaction, key string) (bool, []byte, error) {
	ret := _m.Called(ctx, dbTx, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, string) bool); ok {
		r0 = rf(ctx, dbTx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, string) []byte); ok {
		r1 = rf(ctx, dbTx, key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, database.Transaction, string) error); ok {
		r2 = rf(ctx, dbTx, key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetKey provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) GetKey(_a0 context.Context, _a1 database.Transaction, _a2 *types.AccountIdentifier) (*keys.KeyPair, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *keys.KeyPair
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.AccountIdentifier) *keys.KeyPair); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keys.KeyPair)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction, *types.AccountIdentifier) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Hash provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) Hash(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 string) (*types.TransactionIdentifier, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.TransactionIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, string) *types.TransactionIdentifier); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TransactionIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeadBlockExists provides a mock function with given fields: _a0
func (_m *Helper) HeadBlockExists(_a0 context.Context) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LockedAccounts provides a mock function with given fields: _a0, _a1
func (_m *Helper) LockedAccounts(_a0 context.Context, _a1 database.Transaction) ([]*types.AccountIdentifier, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.AccountIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction) []*types.AccountIdentifier); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.AccountIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Transaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Metadata provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Metadata(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 map[string]interface{}, _a3 []*types.PublicKey) (map[string]interface{}, []*types.Amount, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}, []*types.PublicKey) map[string]interface{}); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 []*types.Amount
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}, []*types.PublicKey) []*types.Amount); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.Amount)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}, []*types.PublicKey) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Parse provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Parse(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 bool, _a3 string) ([]*types.Operation, []*types.AccountIdentifier, map[string]interface{}, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*types.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, bool, string) []*types.Operation); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Operation)
		}
	}

	var r1 []*types.AccountIdentifier
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, bool, string) []*types.AccountIdentifier); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.AccountIdentifier)
		}
	}

	var r2 map[string]interface{}
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, bool, string) map[string]interface{}); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(map[string]interface{})
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, *types.NetworkIdentifier, bool, string) error); ok {
		r3 = rf(_a0, _a1, _a2, _a3)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Payloads provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *Helper) Payloads(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 []*types.Operation, _a3 map[string]interface{}, _a4 []*types.PublicKey) (string, []*types.SigningPayload, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}, []*types.PublicKey) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 []*types.SigningPayload
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}, []*types.PublicKey) []*types.SigningPayload); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.SigningPayload)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}, []*types.PublicKey) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Preprocess provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) Preprocess(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 []*types.Operation, _a3 map[string]interface{}) (map[string]interface{}, []*types.AccountIdentifier, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 []*types.AccountIdentifier
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}) []*types.AccountIdentifier); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.AccountIdentifier)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, []*types.Operation, map[string]interface{}) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetBlob provides a mock function with given fields: ctx, dbTx, key, value
func (_m *Helper) SetBlob(ctx context.Context, dbTx database.Transaction, key string, value []byte) error {
	ret := _m.Called(ctx, dbTx, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, string, []byte) error); ok {
		r0 = rf(ctx, dbTx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sign provides a mock function with given fields: _a0, _a1
func (_m *Helper) Sign(_a0 context.Context, _a1 []*types.SigningPayload) ([]*types.Signature, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.Signature
	if rf, ok := ret.Get(0).(func(context.Context, []*types.SigningPayload) []*types.Signature); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Signature)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*types.SigningPayload) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreKey provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Helper) StoreKey(_a0 context.Context, _a1 database.Transaction, _a2 *types.AccountIdentifier, _a3 *keys.KeyPair) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, *types.AccountIdentifier, *keys.KeyPair) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewHelper creates a new instance of Helper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHelper(t mockConstructorTestingTNewHelper) *Helper {
	mock := &Helper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
