// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockkvstore

import (
	kvstore "github.com/hyperledger-labs/firefly-fabconnect/internal/kvstore"
	mock "github.com/stretchr/testify/mock"
)

// KVStore is an autogenerated mock type for the KVStore type
type KVStore struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *KVStore) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: key
func (_m *KVStore) Delete(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *KVStore) Get(key string) ([]byte, error) {
	ret := _m.Called(key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *KVStore) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIterator provides a mock function with given fields:
func (_m *KVStore) NewIterator() kvstore.KVIterator {
	ret := _m.Called()

	var r0 kvstore.KVIterator
	if rf, ok := ret.Get(0).(func() kvstore.KVIterator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kvstore.KVIterator)
		}
	}

	return r0
}

// NewIteratorWithRange provides a mock function with given fields: keyRange
func (_m *KVStore) NewIteratorWithRange(keyRange interface{}) kvstore.KVIterator {
	ret := _m.Called(keyRange)

	var r0 kvstore.KVIterator
	if rf, ok := ret.Get(0).(func(interface{}) kvstore.KVIterator); ok {
		r0 = rf(keyRange)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kvstore.KVIterator)
		}
	}

	return r0
}

// Put provides a mock function with given fields: key, val
func (_m *KVStore) Put(key string, val []byte) error {
	ret := _m.Called(key, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
