// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/db_interface.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	model "github.com/call-me-snake/products_microservise/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIDb is a mock of IDb interface.
type MockIDb struct {
	ctrl     *gomock.Controller
	recorder *MockIDbMockRecorder
}

// MockIDbMockRecorder is the mock recorder for MockIDb.
type MockIDbMockRecorder struct {
	mock *MockIDb
}

// NewMockIDb creates a new mock instance.
func NewMockIDb(ctrl *gomock.Controller) *MockIDb {
	mock := &MockIDb{ctrl: ctrl}
	mock.recorder = &MockIDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDb) EXPECT() *MockIDbMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIDb) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIDbMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIDb)(nil).Close))
}

// GetProductsList mocks base method.
func (m *MockIDb) GetProductsList() ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsList")
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsList indicates an expected call of GetProductsList.
func (mr *MockIDbMockRecorder) GetProductsList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsList", reflect.TypeOf((*MockIDb)(nil).GetProductsList))
}

// GetProductById mocks base method.
func (m *MockIDb) GetProductById(id int64) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", id)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockIDbMockRecorder) GetProductById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockIDb)(nil).GetProductById), id)
}

// GetProductsByFilter mocks base method.
func (m *MockIDb) GetProductsByFilter(filter map[string]interface{}) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByFilter", filter)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByFilter indicates an expected call of GetProductsByFilter.
func (mr *MockIDbMockRecorder) GetProductsByFilter(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByFilter", reflect.TypeOf((*MockIDb)(nil).GetProductsByFilter), filter)
}

// SetProduct mocks base method.
func (m *MockIDb) SetProduct(newProduct *model.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProduct", newProduct)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProduct indicates an expected call of SetProduct.
func (mr *MockIDbMockRecorder) SetProduct(newProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProduct", reflect.TypeOf((*MockIDb)(nil).SetProduct), newProduct)
}

// UpdateProductById mocks base method.
func (m *MockIDb) UpdateProductById(id int64, params map[string]interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductById", id, params)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductById indicates an expected call of UpdateProductById.
func (mr *MockIDbMockRecorder) UpdateProductById(id, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductById", reflect.TypeOf((*MockIDb)(nil).UpdateProductById), id, params)
}

// DeleteProductById mocks base method.
func (m *MockIDb) DeleteProductById(id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductById", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProductById indicates an expected call of DeleteProductById.
func (mr *MockIDbMockRecorder) DeleteProductById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductById", reflect.TypeOf((*MockIDb)(nil).DeleteProductById), id)
}

// GetCategoriesList mocks base method.
func (m *MockIDb) GetCategoriesList() ([]model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesList")
	ret0, _ := ret[0].([]model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoriesList indicates an expected call of GetCategoriesList.
func (mr *MockIDbMockRecorder) GetCategoriesList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesList", reflect.TypeOf((*MockIDb)(nil).GetCategoriesList))
}

// GetCategoryById mocks base method.
func (m *MockIDb) GetCategoryById(id int64) (*model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryById", id)
	ret0, _ := ret[0].(*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryById indicates an expected call of GetCategoryById.
func (mr *MockIDbMockRecorder) GetCategoryById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryById", reflect.TypeOf((*MockIDb)(nil).GetCategoryById), id)
}

// GetCategoriesByFilter mocks base method.
func (m *MockIDb) GetCategoriesByFilter(filter map[string]interface{}) ([]model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesByFilter", filter)
	ret0, _ := ret[0].([]model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoriesByFilter indicates an expected call of GetCategoriesByFilter.
func (mr *MockIDbMockRecorder) GetCategoriesByFilter(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesByFilter", reflect.TypeOf((*MockIDb)(nil).GetCategoriesByFilter), filter)
}

// SetCategory mocks base method.
func (m *MockIDb) SetCategory(newCategory *model.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCategory", newCategory)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCategory indicates an expected call of SetCategory.
func (mr *MockIDbMockRecorder) SetCategory(newCategory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCategory", reflect.TypeOf((*MockIDb)(nil).SetCategory), newCategory)
}

// UpdateCategoryById mocks base method.
func (m *MockIDb) UpdateCategoryById(id int64, params map[string]interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategoryById", id, params)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategoryById indicates an expected call of UpdateCategoryById.
func (mr *MockIDbMockRecorder) UpdateCategoryById(id, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategoryById", reflect.TypeOf((*MockIDb)(nil).UpdateCategoryById), id, params)
}

// DeleteCategoryById mocks base method.
func (m *MockIDb) DeleteCategoryById(id int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategoryById", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCategoryById indicates an expected call of DeleteCategoryById.
func (mr *MockIDbMockRecorder) DeleteCategoryById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategoryById", reflect.TypeOf((*MockIDb)(nil).DeleteCategoryById), id)
}

// GetUserProductsList mocks base method.
func (m *MockIDb) GetUserProductsList() ([]model.UserProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProductsList")
	ret0, _ := ret[0].([]model.UserProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProductsList indicates an expected call of GetUserProductsList.
func (mr *MockIDbMockRecorder) GetUserProductsList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProductsList", reflect.TypeOf((*MockIDb)(nil).GetUserProductsList))
}

// GetUserProductsListFilteredByCategoryId mocks base method.
func (m *MockIDb) GetUserProductsListFilteredByCategoryId(id int64) ([]model.UserProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProductsListFilteredByCategoryId", id)
	ret0, _ := ret[0].([]model.UserProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProductsListFilteredByCategoryId indicates an expected call of GetUserProductsListFilteredByCategoryId.
func (mr *MockIDbMockRecorder) GetUserProductsListFilteredByCategoryId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProductsListFilteredByCategoryId", reflect.TypeOf((*MockIDb)(nil).GetUserProductsListFilteredByCategoryId), id)
}

// GetUserProductById mocks base method.
func (m *MockIDb) GetUserProductById(id int64) (*model.UserProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProductById", id)
	ret0, _ := ret[0].(*model.UserProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProductById indicates an expected call of GetUserProductById.
func (mr *MockIDbMockRecorder) GetUserProductById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProductById", reflect.TypeOf((*MockIDb)(nil).GetUserProductById), id)
}

// GetProductsByUserFilter mocks base method.
func (m *MockIDb) GetProductsByUserFilter(filter model.FilterUProduct) ([]model.UserProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByUserFilter", filter)
	ret0, _ := ret[0].([]model.UserProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByUserFilter indicates an expected call of GetProductsByUserFilter.
func (mr *MockIDbMockRecorder) GetProductsByUserFilter(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByUserFilter", reflect.TypeOf((*MockIDb)(nil).GetProductsByUserFilter), filter)
}
