// Code generated by MockGen. DO NOT EDIT.
// Source: modules/banner/usecase/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "github.com/azbyluthfan/hello-banner/modules/banner/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockBannerUseCase is a mock of BannerUseCase interface
type MockBannerUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockBannerUseCaseMockRecorder
}

// MockBannerUseCaseMockRecorder is the mock recorder for MockBannerUseCase
type MockBannerUseCaseMockRecorder struct {
	mock *MockBannerUseCase
}

// NewMockBannerUseCase creates a new mock instance
func NewMockBannerUseCase(ctrl *gomock.Controller) *MockBannerUseCase {
	mock := &MockBannerUseCase{ctrl: ctrl}
	mock.recorder = &MockBannerUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBannerUseCase) EXPECT() *MockBannerUseCaseMockRecorder {
	return m.recorder
}

// GetBanner mocks base method
func (m *MockBannerUseCase) GetBanner(time time.Time, timezone, ipAddress string) (*model.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanner", time, timezone, ipAddress)
	ret0, _ := ret[0].(*model.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBanner indicates an expected call of GetBanner
func (mr *MockBannerUseCaseMockRecorder) GetBanner(time, timezone, ipAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanner", reflect.TypeOf((*MockBannerUseCase)(nil).GetBanner), time, timezone, ipAddress)
}
