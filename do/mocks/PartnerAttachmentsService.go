// Code generated by MockGen. DO NOT EDIT.
// Source: partner_network_connect.go
//
// Generated by this command:
//
//	mockgen -source partner_network_connect.go -package=mocks PartnerAttachmentsService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	do "github.com/digitalocean/doctl/do"
	godo "github.com/digitalocean/godo"
	gomock "go.uber.org/mock/gomock"
)

// MockPartnerAttachmentsService is a mock of PartnerAttachmentsService interface.
type MockPartnerAttachmentsService struct {
	ctrl     *gomock.Controller
	recorder *MockPartnerAttachmentsServiceMockRecorder
	isgomock struct{}
}

// MockPartnerAttachmentsServiceMockRecorder is the mock recorder for MockPartnerAttachmentsService.
type MockPartnerAttachmentsServiceMockRecorder struct {
	mock *MockPartnerAttachmentsService
}

// NewMockPartnerAttachmentsService creates a new mock instance.
func NewMockPartnerAttachmentsService(ctrl *gomock.Controller) *MockPartnerAttachmentsService {
	mock := &MockPartnerAttachmentsService{ctrl: ctrl}
	mock.recorder = &MockPartnerAttachmentsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPartnerAttachmentsService) EXPECT() *MockPartnerAttachmentsServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPartnerAttachmentsService) Create(arg0 *godo.PartnerAttachmentCreateRequest) (*do.PartnerAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*do.PartnerAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPartnerAttachmentsServiceMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).Create), arg0)
}

// DeletePartnerAttachment mocks base method.
func (m *MockPartnerAttachmentsService) DeletePartnerAttachment(paID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePartnerAttachment", paID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePartnerAttachment indicates an expected call of DeletePartnerAttachment.
func (mr *MockPartnerAttachmentsServiceMockRecorder) DeletePartnerAttachment(paID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePartnerAttachment", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).DeletePartnerAttachment), paID)
}

// GetBGPAuthKey mocks base method.
func (m *MockPartnerAttachmentsService) GetBGPAuthKey(paID string) (*do.PartnerAttachmentBGPAuthKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBGPAuthKey", paID)
	ret0, _ := ret[0].(*do.PartnerAttachmentBGPAuthKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBGPAuthKey indicates an expected call of GetBGPAuthKey.
func (mr *MockPartnerAttachmentsServiceMockRecorder) GetBGPAuthKey(paID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBGPAuthKey", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).GetBGPAuthKey), paID)
}

// GetPartnerAttachment mocks base method.
func (m *MockPartnerAttachmentsService) GetPartnerAttachment(paID string) (*do.PartnerAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartnerAttachment", paID)
	ret0, _ := ret[0].(*do.PartnerAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartnerAttachment indicates an expected call of GetPartnerAttachment.
func (mr *MockPartnerAttachmentsServiceMockRecorder) GetPartnerAttachment(paID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartnerAttachment", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).GetPartnerAttachment), paID)
}

// GetServiceKey mocks base method.
func (m *MockPartnerAttachmentsService) GetServiceKey(paID string) (*do.PartnerAttachmentServiceKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceKey", paID)
	ret0, _ := ret[0].(*do.PartnerAttachmentServiceKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceKey indicates an expected call of GetServiceKey.
func (mr *MockPartnerAttachmentsServiceMockRecorder) GetServiceKey(paID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceKey", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).GetServiceKey), paID)
}

// ListPartnerAttachmentRoutes mocks base method.
func (m *MockPartnerAttachmentsService) ListPartnerAttachmentRoutes(paID string) (do.PartnerAttachmentRoutes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPartnerAttachmentRoutes", paID)
	ret0, _ := ret[0].(do.PartnerAttachmentRoutes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPartnerAttachmentRoutes indicates an expected call of ListPartnerAttachmentRoutes.
func (mr *MockPartnerAttachmentsServiceMockRecorder) ListPartnerAttachmentRoutes(paID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPartnerAttachmentRoutes", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).ListPartnerAttachmentRoutes), paID)
}

// ListPartnerAttachments mocks base method.
func (m *MockPartnerAttachmentsService) ListPartnerAttachments() (do.PartnerAttachments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPartnerAttachments")
	ret0, _ := ret[0].(do.PartnerAttachments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPartnerAttachments indicates an expected call of ListPartnerAttachments.
func (mr *MockPartnerAttachmentsServiceMockRecorder) ListPartnerAttachments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPartnerAttachments", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).ListPartnerAttachments))
}

// RegenerateServiceKey mocks base method.
func (m *MockPartnerAttachmentsService) RegenerateServiceKey(paID string) (*do.PartnerAttachmentRegenerateServiceKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegenerateServiceKey", paID)
	ret0, _ := ret[0].(*do.PartnerAttachmentRegenerateServiceKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegenerateServiceKey indicates an expected call of RegenerateServiceKey.
func (mr *MockPartnerAttachmentsServiceMockRecorder) RegenerateServiceKey(paID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegenerateServiceKey", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).RegenerateServiceKey), paID)
}

// UpdatePartnerAttachment mocks base method.
func (m *MockPartnerAttachmentsService) UpdatePartnerAttachment(paID string, req *godo.PartnerAttachmentUpdateRequest) (*do.PartnerAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePartnerAttachment", paID, req)
	ret0, _ := ret[0].(*do.PartnerAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePartnerAttachment indicates an expected call of UpdatePartnerAttachment.
func (mr *MockPartnerAttachmentsServiceMockRecorder) UpdatePartnerAttachment(paID, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePartnerAttachment", reflect.TypeOf((*MockPartnerAttachmentsService)(nil).UpdatePartnerAttachment), paID, req)
}
