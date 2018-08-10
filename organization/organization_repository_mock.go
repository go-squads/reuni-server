// Code generated by MockGen. DO NOT EDIT.
// Source: organization/organization_repository.go

// Package organization is a generated GoMock package.
package organization

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockrepository is a mock of repository interface
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
}

// MockrepositoryMockRecorder is the mock recorder for Mockrepository
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// NewMockrepository creates a new mock instance
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// createNewOrganization mocks base method
func (m *Mockrepository) createNewOrganization(organization_name string) (int64, error) {
	ret := m.ctrl.Call(m, "createNewOrganization", organization_name)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// createNewOrganization indicates an expected call of createNewOrganization
func (mr *MockrepositoryMockRecorder) createNewOrganization(organization_name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createNewOrganization", reflect.TypeOf((*Mockrepository)(nil).createNewOrganization), organization_name)
}

// addUser mocks base method
func (m *Mockrepository) addUser(organizationId, userId int64, role string) error {
	ret := m.ctrl.Call(m, "addUser", organizationId, userId, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// addUser indicates an expected call of addUser
func (mr *MockrepositoryMockRecorder) addUser(organizationId, userId, role interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addUser", reflect.TypeOf((*Mockrepository)(nil).addUser), organizationId, userId, role)
}

// deleteUser mocks base method
func (m *Mockrepository) deleteUser(organizationId, userId int64) error {
	ret := m.ctrl.Call(m, "deleteUser", organizationId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// deleteUser indicates an expected call of deleteUser
func (mr *MockrepositoryMockRecorder) deleteUser(organizationId, userId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteUser", reflect.TypeOf((*Mockrepository)(nil).deleteUser), organizationId, userId)
}

// updateRoleOfUser mocks base method
func (m *Mockrepository) updateRoleOfUser(newRole string, organizationId, userId int64) error {
	ret := m.ctrl.Call(m, "updateRoleOfUser", newRole, organizationId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// updateRoleOfUser indicates an expected call of updateRoleOfUser
func (mr *MockrepositoryMockRecorder) updateRoleOfUser(newRole, organizationId, userId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateRoleOfUser", reflect.TypeOf((*Mockrepository)(nil).updateRoleOfUser), newRole, organizationId, userId)
}

// getAllMemberOfOrganization mocks base method
func (m *Mockrepository) getAllMemberOfOrganization(organizationId int64) ([]map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "getAllMemberOfOrganization", organizationId)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getAllMemberOfOrganization indicates an expected call of getAllMemberOfOrganization
func (mr *MockrepositoryMockRecorder) getAllMemberOfOrganization(organizationId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAllMemberOfOrganization", reflect.TypeOf((*Mockrepository)(nil).getAllMemberOfOrganization), organizationId)
}

// getAllOrganization mocks base method
func (m *Mockrepository) getAllOrganization(userId int) ([]OrganizationMember, error) {
	ret := m.ctrl.Call(m, "getAllOrganization", userId)
	ret0, _ := ret[0].([]OrganizationMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getAllOrganization indicates an expected call of getAllOrganization
func (mr *MockrepositoryMockRecorder) getAllOrganization(userId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAllOrganization", reflect.TypeOf((*Mockrepository)(nil).getAllOrganization), userId)
}
