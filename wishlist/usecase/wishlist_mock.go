// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	models "go-clean-arch-test/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateWish mocks base method.
func (m *MockRepository) CreateWish(ctx context.Context, wish *models.Wish) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWish", ctx, wish)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWish indicates an expected call of CreateWish.
func (mr *MockRepositoryMockRecorder) CreateWish(ctx, wish interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWish", reflect.TypeOf((*MockRepository)(nil).CreateWish), ctx, wish)
}

// DeleteWishByID mocks base method.
func (m *MockRepository) DeleteWishByID(ctx context.Context, wish *models.Wish) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWishByID", ctx, wish)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWishByID indicates an expected call of DeleteWishByID.
func (mr *MockRepositoryMockRecorder) DeleteWishByID(ctx, wish interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWishByID", reflect.TypeOf((*MockRepository)(nil).DeleteWishByID), ctx, wish)
}

// GetAllWishes mocks base method.
func (m *MockRepository) GetAllWishes(ctx context.Context) ([]*models.Wish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWishes", ctx)
	ret0, _ := ret[0].([]*models.Wish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWishes indicates an expected call of GetAllWishes.
func (mr *MockRepositoryMockRecorder) GetAllWishes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWishes", reflect.TypeOf((*MockRepository)(nil).GetAllWishes), ctx)
}