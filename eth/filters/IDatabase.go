// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ethereum/go-ethereum/ethdb (interfaces: Database)

// Package filters is a generated GoMock package.
package filters

import (
	reflect "reflect"

	ethdb "github.com/ethereum/go-ethereum/ethdb"
	gomock "github.com/golang/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Ancient mocks base method.
func (m *MockDatabase) Ancient(arg0 string, arg1 uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ancient", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ancient indicates an expected call of Ancient.
func (mr *MockDatabaseMockRecorder) Ancient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ancient", reflect.TypeOf((*MockDatabase)(nil).Ancient), arg0, arg1)
}

// AncientDatadir mocks base method.
func (m *MockDatabase) AncientDatadir() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AncientDatadir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AncientDatadir indicates an expected call of AncientDatadir.
func (mr *MockDatabaseMockRecorder) AncientDatadir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AncientDatadir", reflect.TypeOf((*MockDatabase)(nil).AncientDatadir))
}

// AncientRange mocks base method.
func (m *MockDatabase) AncientRange(arg0 string, arg1, arg2, arg3 uint64) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AncientRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AncientRange indicates an expected call of AncientRange.
func (mr *MockDatabaseMockRecorder) AncientRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AncientRange", reflect.TypeOf((*MockDatabase)(nil).AncientRange), arg0, arg1, arg2, arg3)
}

// AncientSize mocks base method.
func (m *MockDatabase) AncientSize(arg0 string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AncientSize", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AncientSize indicates an expected call of AncientSize.
func (mr *MockDatabaseMockRecorder) AncientSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AncientSize", reflect.TypeOf((*MockDatabase)(nil).AncientSize), arg0)
}

// Ancients mocks base method.
func (m *MockDatabase) Ancients() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ancients")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ancients indicates an expected call of Ancients.
func (mr *MockDatabaseMockRecorder) Ancients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ancients", reflect.TypeOf((*MockDatabase)(nil).Ancients))
}

// Close mocks base method.
func (m *MockDatabase) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatabaseMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDatabase)(nil).Close))
}

// Compact mocks base method.
func (m *MockDatabase) Compact(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compact", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Compact indicates an expected call of Compact.
func (mr *MockDatabaseMockRecorder) Compact(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compact", reflect.TypeOf((*MockDatabase)(nil).Compact), arg0, arg1)
}

// Delete mocks base method.
func (m *MockDatabase) Delete(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDatabaseMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatabase)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockDatabase) Get(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDatabaseMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatabase)(nil).Get), arg0)
}

// Has mocks base method.
func (m *MockDatabase) Has(arg0 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has.
func (mr *MockDatabaseMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockDatabase)(nil).Has), arg0)
}

// HasAncient mocks base method.
func (m *MockDatabase) HasAncient(arg0 string, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAncient", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasAncient indicates an expected call of HasAncient.
func (mr *MockDatabaseMockRecorder) HasAncient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAncient", reflect.TypeOf((*MockDatabase)(nil).HasAncient), arg0, arg1)
}

// MigrateTable mocks base method.
func (m *MockDatabase) MigrateTable(arg0 string, arg1 func([]byte) ([]byte, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateTable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateTable indicates an expected call of MigrateTable.
func (mr *MockDatabaseMockRecorder) MigrateTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateTable", reflect.TypeOf((*MockDatabase)(nil).MigrateTable), arg0, arg1)
}

// ModifyAncients mocks base method.
func (m *MockDatabase) ModifyAncients(arg0 func(ethdb.AncientWriteOp) error) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyAncients", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyAncients indicates an expected call of ModifyAncients.
func (mr *MockDatabaseMockRecorder) ModifyAncients(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyAncients", reflect.TypeOf((*MockDatabase)(nil).ModifyAncients), arg0)
}

// NewBatch mocks base method.
func (m *MockDatabase) NewBatch() ethdb.Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBatch")
	ret0, _ := ret[0].(ethdb.Batch)
	return ret0
}

// NewBatch indicates an expected call of NewBatch.
func (mr *MockDatabaseMockRecorder) NewBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBatch", reflect.TypeOf((*MockDatabase)(nil).NewBatch))
}

// NewBatchWithSize mocks base method.
func (m *MockDatabase) NewBatchWithSize(arg0 int) ethdb.Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBatchWithSize", arg0)
	ret0, _ := ret[0].(ethdb.Batch)
	return ret0
}

// NewBatchWithSize indicates an expected call of NewBatchWithSize.
func (mr *MockDatabaseMockRecorder) NewBatchWithSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBatchWithSize", reflect.TypeOf((*MockDatabase)(nil).NewBatchWithSize), arg0)
}

// NewIterator mocks base method.
func (m *MockDatabase) NewIterator(arg0, arg1 []byte) ethdb.Iterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewIterator", arg0, arg1)
	ret0, _ := ret[0].(ethdb.Iterator)
	return ret0
}

// NewIterator indicates an expected call of NewIterator.
func (mr *MockDatabaseMockRecorder) NewIterator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIterator", reflect.TypeOf((*MockDatabase)(nil).NewIterator), arg0, arg1)
}

// NewSnapshot mocks base method.
func (m *MockDatabase) NewSnapshot() (ethdb.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSnapshot")
	ret0, _ := ret[0].(ethdb.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSnapshot indicates an expected call of NewSnapshot.
func (mr *MockDatabaseMockRecorder) NewSnapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSnapshot", reflect.TypeOf((*MockDatabase)(nil).NewSnapshot))
}

// Put mocks base method.
func (m *MockDatabase) Put(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockDatabaseMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockDatabase)(nil).Put), arg0, arg1)
}

// ReadAncients mocks base method.
func (m *MockDatabase) ReadAncients(arg0 func(ethdb.AncientReaderOp) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAncients", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadAncients indicates an expected call of ReadAncients.
func (mr *MockDatabaseMockRecorder) ReadAncients(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAncients", reflect.TypeOf((*MockDatabase)(nil).ReadAncients), arg0)
}

// Stat mocks base method.
func (m *MockDatabase) Stat(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockDatabaseMockRecorder) Stat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockDatabase)(nil).Stat), arg0)
}

// Sync mocks base method.
func (m *MockDatabase) Sync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockDatabaseMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockDatabase)(nil).Sync))
}

// Tail mocks base method.
func (m *MockDatabase) Tail() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tail")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tail indicates an expected call of Tail.
func (mr *MockDatabaseMockRecorder) Tail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tail", reflect.TypeOf((*MockDatabase)(nil).Tail))
}

// TruncateHead mocks base method.
func (m *MockDatabase) TruncateHead(arg0 uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TruncateHead", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TruncateHead indicates an expected call of TruncateHead.
func (mr *MockDatabaseMockRecorder) TruncateHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateHead", reflect.TypeOf((*MockDatabase)(nil).TruncateHead), arg0)
}

// TruncateTail mocks base method.
func (m *MockDatabase) TruncateTail(arg0 uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TruncateTail", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TruncateTail indicates an expected call of TruncateTail.
func (mr *MockDatabaseMockRecorder) TruncateTail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateTail", reflect.TypeOf((*MockDatabase)(nil).TruncateTail), arg0)
}
