/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package outofband

import (
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/outofband"
)

// MockOobService is a mock of OobService interface.
type MockOobService struct {
	AcceptInvitationHandle      func(*outofband.Invitation, string, []string) (string, error)
	AcceptRequestHandle         func(*outofband.Request, string, []string) (string, error)
	ActionContinueHandle        func(string, outofband.Options) error
	ActionStopHandle            func(string, error) error
	ActionsHandle               func() ([]outofband.Action, error)
	RegisterActionEventHandle   func(chan<- service.DIDCommAction) error
	RegisterMsgEventHandle      func(chan<- service.StateMsg) error
	SaveInvitationHandle        func(*outofband.Invitation) error
	SaveRequestHandle           func(*outofband.Request) error
	UnregisterActionEventHandle func(chan<- service.DIDCommAction) error
	UnregisterMsgEventHandle    func(chan<- service.StateMsg) error
}

// AcceptInvitation mock implementation.
func (m *MockOobService) AcceptInvitation(arg0 *outofband.Invitation, arg1 string, arg2 []string) (string, error) {
	if m.AcceptInvitationHandle != nil {
		return m.AcceptInvitationHandle(arg0, arg1, arg2)
	}

	return "", nil
}

// AcceptRequest mock implementation.
func (m *MockOobService) AcceptRequest(arg0 *outofband.Request, arg1 string, arg2 []string) (string, error) {
	if m.AcceptRequestHandle != nil {
		return m.AcceptRequestHandle(arg0, arg1, arg2)
	}

	return "", nil
}

// ActionContinue mock implementation.
func (m *MockOobService) ActionContinue(arg0 string, arg1 outofband.Options) error {
	if m.ActionContinueHandle != nil {
		return m.ActionContinueHandle(arg0, arg1)
	}

	return nil
}

// ActionStop mock implementation.
func (m *MockOobService) ActionStop(arg0 string, arg1 error) error {
	if m.ActionStopHandle != nil {
		return m.ActionStopHandle(arg0, arg1)
	}

	return nil
}

// Actions mock implementation.
func (m *MockOobService) Actions() ([]outofband.Action, error) {
	if m.ActionsHandle != nil {
		return m.ActionsHandle()
	}

	return []outofband.Action{}, nil
}

// RegisterActionEvent mock implementation.
func (m *MockOobService) RegisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	if m.RegisterActionEventHandle != nil {
		return m.RegisterActionEventHandle(arg0)
	}

	return nil
}

// RegisterMsgEvent mock implementation.
func (m *MockOobService) RegisterMsgEvent(arg0 chan<- service.StateMsg) error {
	if m.RegisterMsgEventHandle != nil {
		return m.RegisterMsgEventHandle(arg0)
	}

	return nil
}

// SaveInvitation mock implementation.
func (m *MockOobService) SaveInvitation(arg0 *outofband.Invitation) error {
	if m.SaveInvitationHandle != nil {
		return m.SaveInvitationHandle(arg0)
	}

	return nil
}

// SaveRequest mock implementation.
func (m *MockOobService) SaveRequest(arg0 *outofband.Request) error {
	if m.SaveRequestHandle != nil {
		return m.SaveRequestHandle(arg0)
	}

	return nil
}

// UnregisterActionEvent mock implementation.
func (m *MockOobService) UnregisterActionEvent(arg0 chan<- service.DIDCommAction) error {
	if m.UnregisterActionEventHandle != nil {
		return m.UnregisterActionEventHandle(arg0)
	}

	return nil
}

// UnregisterMsgEvent mock implementation.
func (m *MockOobService) UnregisterMsgEvent(arg0 chan<- service.StateMsg) error {
	if m.UnregisterMsgEventHandle != nil {
		return m.UnregisterMsgEventHandle(arg0)
	}

	return nil
}
