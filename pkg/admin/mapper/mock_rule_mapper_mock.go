/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apache/dubbo-admin/pkg/admin/mapper (interfaces: MockRuleMapper)

// Package mapper is a generated GoMock package.
package mapper

import (
	context "context"
	reflect "reflect"

	model "github.com/apache/dubbo-admin/pkg/admin/model"
	gomock "github.com/golang/mock/gomock"
)

// MockMockRuleMapper is a mock of MockRuleMapper interface.
type MockMockRuleMapper struct {
	ctrl     *gomock.Controller
	recorder *MockMockRuleMapperMockRecorder
}

// MockMockRuleMapperMockRecorder is the mock recorder for MockMockRuleMapper.
type MockMockRuleMapperMockRecorder struct {
	mock *MockMockRuleMapper
}

// NewMockMockRuleMapper creates a new mock instance.
func NewMockMockRuleMapper(ctrl *gomock.Controller) *MockMockRuleMapper {
	mock := &MockMockRuleMapper{ctrl: ctrl}
	mock.recorder = &MockMockRuleMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMockRuleMapper) EXPECT() *MockMockRuleMapperMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMockRuleMapper) Create(arg0 *model.MockRuleEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMockRuleMapperMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMockRuleMapper)(nil).Create), arg0)
}

// DeleteById mocks base method.
func (m *MockMockRuleMapper) DeleteById(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockMockRuleMapperMockRecorder) DeleteById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockMockRuleMapper)(nil).DeleteById), arg0)
}

// FindByPage mocks base method.
func (m *MockMockRuleMapper) FindByPage(arg0 string, arg1, arg2 int) ([]*model.MockRuleEntity, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPage", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.MockRuleEntity)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindByPage indicates an expected call of FindByPage.
func (mr *MockMockRuleMapperMockRecorder) FindByPage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPage", reflect.TypeOf((*MockMockRuleMapper)(nil).FindByPage), arg0, arg1, arg2)
}

// FindByServiceNameAndMethodName mocks base method.
func (m *MockMockRuleMapper) FindByServiceNameAndMethodName(arg0 context.Context, arg1, arg2 string) (*model.MockRuleEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByServiceNameAndMethodName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.MockRuleEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByServiceNameAndMethodName indicates an expected call of FindByServiceNameAndMethodName.
func (mr *MockMockRuleMapperMockRecorder) FindByServiceNameAndMethodName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByServiceNameAndMethodName", reflect.TypeOf((*MockMockRuleMapper)(nil).FindByServiceNameAndMethodName), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockMockRuleMapper) Update(arg0 *model.MockRuleEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMockRuleMapperMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMockRuleMapper)(nil).Update), arg0)
}
