// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	parking "github.com/Lucifer07/parking-lot/parking"
	mock "github.com/stretchr/testify/mock"
)

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

// NotifyObserver provides a mock function with given fields: condition
func (_m *Publisher) NotifyObserver(condition string) {
	_m.Called(condition)
}

// Register provides a mock function with given fields: _a0
func (_m *Publisher) Register(_a0 parking.Observer) {
	_m.Called(_a0)
}
