// Code generated by MockGen. DO NOT EDIT.
// Source: s3runtime/types.go

// Package s3_mocks is a generated GoMock package.
package s3_mocks

import (
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/yuvalman/s3BucketController/api/v1"
)

// MockS3Client is a mock of S3Client interface.
type MockS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientMockRecorder
}

// MockS3ClientMockRecorder is the mock recorder for MockS3Client.
type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

// NewMockS3Client creates a new mock instance.
func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

// PutPublicAccessBlock mocks base method.
func (m *MockS3Client) PutPublicAccessBlock(ctx context.Context, params *s3.PutPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.PutPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutPublicAccessBlock", varargs...)
	ret0, _ := ret[0].(*s3.PutPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutPublicAccessBlock indicates an expected call of PutPublicAccessBlock.
func (mr *MockS3ClientMockRecorder) PutPublicAccessBlock(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPublicAccessBlock", reflect.TypeOf((*MockS3Client)(nil).PutPublicAccessBlock), varargs...)
}

// MockS3ops is a mock of S3ops interface.
type MockS3ops struct {
	ctrl     *gomock.Controller
	recorder *MockS3opsMockRecorder
}

// MockS3opsMockRecorder is the mock recorder for MockS3ops.
type MockS3opsMockRecorder struct {
	mock *MockS3ops
}

// NewMockS3ops creates a new mock instance.
func NewMockS3ops(ctrl *gomock.Controller) *MockS3ops {
	mock := &MockS3ops{ctrl: ctrl}
	mock.recorder = &MockS3opsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3ops) EXPECT() *MockS3opsMockRecorder {
	return m.recorder
}

// UpdatePublicAccessBlock mocks base method.
func (m *MockS3ops) UpdatePublicAccessBlock(ctx context.Context, bucket *v1.S3Bucket, log logr.Logger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePublicAccessBlock", ctx, bucket, log)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePublicAccessBlock indicates an expected call of UpdatePublicAccessBlock.
func (mr *MockS3opsMockRecorder) UpdatePublicAccessBlock(ctx, bucket, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePublicAccessBlock", reflect.TypeOf((*MockS3ops)(nil).UpdatePublicAccessBlock), ctx, bucket, log)
}