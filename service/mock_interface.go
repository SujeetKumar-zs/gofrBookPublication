// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package service is a generated GoMock package.
package service

import (
	entities "Project/entities"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBook is a mock of Book interface.
type MockBook struct {
	ctrl     *gomock.Controller
	recorder *MockBookMockRecorder
}

// MockBookMockRecorder is the mock recorder for MockBook.
type MockBookMockRecorder struct {
	mock *MockBook
}

// NewMockBook creates a new mock instance.
func NewMockBook(ctrl *gomock.Controller) *MockBook {
	mock := &MockBook{ctrl: ctrl}
	mock.recorder = &MockBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBook) EXPECT() *MockBookMockRecorder {
	return m.recorder
}

// DeleteBook mocks base method.
func (m *MockBook) DeleteBook(ctx context.Context, id int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", ctx, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockBookMockRecorder) DeleteBook(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockBook)(nil).DeleteBook), ctx, id)
}

// GetAllBook mocks base method.
func (m *MockBook) GetAllBook(ctx context.Context, title, includeAuthor string) ([]entities.Books, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBook", ctx, title, includeAuthor)
	ret0, _ := ret[0].([]entities.Books)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBook indicates an expected call of GetAllBook.
func (mr *MockBookMockRecorder) GetAllBook(ctx, title, includeAuthor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBook", reflect.TypeOf((*MockBook)(nil).GetAllBook), ctx, title, includeAuthor)
}

// GetBookByID mocks base method.
func (m *MockBook) GetBookByID(ctx context.Context, id int) (entities.Books, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", ctx, id)
	ret0, _ := ret[0].(entities.Books)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockBookMockRecorder) GetBookByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockBook)(nil).GetBookByID), ctx, id)
}

// PostBook mocks base method.
func (m *MockBook) PostBook(ctx context.Context, b entities.Books) (entities.Books, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBook", ctx, b)
	ret0, _ := ret[0].(entities.Books)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBook indicates an expected call of PostBook.
func (mr *MockBookMockRecorder) PostBook(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBook", reflect.TypeOf((*MockBook)(nil).PostBook), ctx, b)
}

// PutBook mocks base method.
func (m *MockBook) PutBook(ctx context.Context, id int, b entities.Books) (entities.Books, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutBook", ctx, id, b)
	ret0, _ := ret[0].(entities.Books)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBook indicates an expected call of PutBook.
func (mr *MockBookMockRecorder) PutBook(ctx, id, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBook", reflect.TypeOf((*MockBook)(nil).PutBook), ctx, id, b)
}

// MockAuthor is a mock of Author interface.
type MockAuthor struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorMockRecorder
}

// MockAuthorMockRecorder is the mock recorder for MockAuthor.
type MockAuthorMockRecorder struct {
	mock *MockAuthor
}

// NewMockAuthor creates a new mock instance.
func NewMockAuthor(ctrl *gomock.Controller) *MockAuthor {
	mock := &MockAuthor{ctrl: ctrl}
	mock.recorder = &MockAuthorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthor) EXPECT() *MockAuthorMockRecorder {
	return m.recorder
}

// DeleteAuthor mocks base method.
func (m *MockAuthor) DeleteAuthor(ctx context.Context, id int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthor", ctx, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAuthor indicates an expected call of DeleteAuthor.
func (mr *MockAuthorMockRecorder) DeleteAuthor(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthor", reflect.TypeOf((*MockAuthor)(nil).DeleteAuthor), ctx, id)
}

// PostAuthor mocks base method.
func (m *MockAuthor) PostAuthor(ctx context.Context, a entities.Author) (entities.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostAuthor", ctx, a)
	ret0, _ := ret[0].(entities.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostAuthor indicates an expected call of PostAuthor.
func (mr *MockAuthorMockRecorder) PostAuthor(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostAuthor", reflect.TypeOf((*MockAuthor)(nil).PostAuthor), ctx, a)
}

// PutAuthor mocks base method.
func (m *MockAuthor) PutAuthor(ctx context.Context, a entities.Author) (entities.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutAuthor", ctx, a)
	ret0, _ := ret[0].(entities.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutAuthor indicates an expected call of PutAuthor.
func (mr *MockAuthorMockRecorder) PutAuthor(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutAuthor", reflect.TypeOf((*MockAuthor)(nil).PutAuthor), ctx, a)
}
