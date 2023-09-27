// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

package grpc_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/0x6flab/mpesaoverlay/grpc"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa"
	"github.com/0x6flab/mpesaoverlay/pkg/mpesa/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	errMock = errors.New("mock error")

	validResp = mpesa.ValidResp{
		OriginatorConversationID: "AG_20230907_2010325b025970fde878",
		ConversationID:           "AG_20230907_2010325b025970fde878",
		ResponseDescription:      "Accept the service request successfully.",
		ResponseCode:             "0",
	}
)

func TestToken(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		expectedResp mpesa.TokenResp
		expectedErr  error
	}{
		{
			name: "Token success",
			expectedResp: mpesa.TokenResp{
				AccessToken: "mocked-token",
				Expiry:      "3559",
			},
			expectedErr: nil,
		},
		{
			name:         "Token error",
			expectedResp: mpesa.TokenResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("Token").Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.Token()

		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))
		assert.Equal(t, tc.expectedErr, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))

		call.Parent.AssertCalled(t, "Token")
		call.Unset()
	}
}

func TestAccountBalance(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.AccountBalanceReq
		expectedResp mpesa.AccountBalanceResp
		expectedErr  error
	}{
		{
			name: "AccountBalance success",
			req:  mpesa.AccountBalanceReq{},
			expectedResp: mpesa.AccountBalanceResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:         "AccountBalance error",
			req:          mpesa.AccountBalanceReq{},
			expectedResp: mpesa.AccountBalanceResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("AccountBalance", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.AccountBalance(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestC2BRegisterURL(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.C2BRegisterURLReq
		expectedResp mpesa.C2BRegisterURLResp
		expectedErr  error
	}{
		{
			name: "C2BRegisterURL success",
			req:  mpesa.C2BRegisterURLReq{},
			expectedResp: mpesa.C2BRegisterURLResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:         "C2BRegisterURL error",
			req:          mpesa.C2BRegisterURLReq{},
			expectedResp: mpesa.C2BRegisterURLResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("C2BRegisterURL", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.C2BRegisterURL(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestC2BSimulate(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.C2BSimulateReq
		expectedResp mpesa.C2BSimulateResp
		expectedErr  error
	}{
		{
			name: "C2BSimulate success",
			req:  mpesa.C2BSimulateReq{},
			expectedResp: mpesa.C2BSimulateResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:         "C2BSimulate error",
			req:          mpesa.C2BSimulateReq{},
			expectedResp: mpesa.C2BSimulateResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("C2BSimulate", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.C2BSimulate(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestGenerateQR(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.GenerateQRReq
		expectedResp mpesa.GenerateQRResp
		expectedErr  error
	}{
		{
			name: "GenerateQR success",
			req:  mpesa.GenerateQRReq{},
			expectedResp: mpesa.GenerateQRResp{
				ResponseDescription: "The service request is processed successfully.",
				ResponseCode:        "00",
				RequestID:           "QRCode:...",
				QRCode:              "qr_code",
			},
			expectedErr: nil,
		},
		{
			name:         "GenerateQR error",
			req:          mpesa.GenerateQRReq{},
			expectedResp: mpesa.GenerateQRResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("GenerateQR", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.GenerateQR(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestExpressQuery(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.ExpressQueryReq
		expectedResp mpesa.ExpressQueryResp
		expectedErr  error
	}{
		{
			name: "ExpressQuery success",
			req:  mpesa.ExpressQueryReq{},
			expectedResp: mpesa.ExpressQueryResp{
				ResponseDescription: "The service request has been accepted successsfully",
				ResponseCode:        "0",
				MerchantRequestID:   "92643-47073138-2",
				CheckoutRequestID:   "ws_CO_07092023195244460712345678",
				CustomerMessage:     "",
				ResultCode:          "1032",
				ResultDesc:          "Request cancelled by user",
			},
			expectedErr: nil,
		},
		{
			name:         "ExpressQuery error",
			req:          mpesa.ExpressQueryReq{},
			expectedResp: mpesa.ExpressQueryResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("ExpressQuery", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.ExpressQuery(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestReverse(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.ReverseReq
		expectedResp mpesa.ReverseResp
		expectedErr  error
	}{
		{
			name: "Reverse success",
			req:  mpesa.ReverseReq{},
			expectedResp: mpesa.ReverseResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:         "Reverse error",
			req:          mpesa.ReverseReq{},
			expectedResp: mpesa.ReverseResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("Reverse", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.Reverse(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestExpressSimulate(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.ExpressSimulateReq
		expectedResp mpesa.ExpressSimulateResp
		expectedErr  error
	}{
		{
			name: "ExpressSimulate success",
			req:  mpesa.ExpressSimulateReq{},
			expectedResp: mpesa.ExpressSimulateResp{
				ResponseDescription: "Success. Request accepted for processing",
				ResponseCode:        "0",
				MerchantRequestID:   "27260-79456854-2",
				CheckoutRequestID:   "ws_CO_07092023004130971712345678",
				CustomerMessage:     "Success. Request accepted for processing",
			},
			expectedErr: nil,
		},
		{
			name:         "ExpressSimulate error",
			req:          mpesa.ExpressSimulateReq{},
			expectedResp: mpesa.ExpressSimulateResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("ExpressSimulate", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.ExpressSimulate(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestRemitTax(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.RemitTaxReq
		expectedResp mpesa.RemitTaxResp
		expectedErr  error
	}{
		{
			name: "RemitTax success",
			req:  mpesa.RemitTaxReq{},
			expectedResp: mpesa.RemitTaxResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:         "RemitTax error",
			req:          mpesa.RemitTaxReq{},
			expectedResp: mpesa.RemitTaxResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("RemitTax", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.RemitTax(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestTransactionStatus(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.TransactionStatusReq
		expectedResp mpesa.TransactionStatusResp
		expectedErr  error
	}{
		{
			name: "TransactionStatus success",
			req:  mpesa.TransactionStatusReq{},
			expectedResp: mpesa.TransactionStatusResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:         "TransactionStatus error",
			req:          mpesa.TransactionStatusReq{},
			expectedResp: mpesa.TransactionStatusResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("TransactionStatus", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.TransactionStatus(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}

func TestB2CPayment(t *testing.T) {
	mockSDK := new(mocks.SDK)
	s := grpc.NewService(mockSDK)

	cases := []struct {
		name         string
		req          mpesa.B2CPaymentReq
		expectedResp mpesa.B2CPaymentResp
		expectedErr  error
	}{
		{
			name: "B2CPayment success",
			req:  mpesa.B2CPaymentReq{},
			expectedResp: mpesa.B2CPaymentResp{
				ValidResp: validResp,
			},
			expectedErr: nil,
		},
		{
			name:         "B2CPayment error",
			req:          mpesa.B2CPaymentReq{},
			expectedResp: mpesa.B2CPaymentResp{},
			expectedErr:  errMock,
		},
	}

	for _, tc := range cases {
		call := mockSDK.On("B2CPayment", mock.Anything).Return(tc.expectedResp, tc.expectedErr)

		resp, err := s.B2CPayment(tc.req)
		if err != nil {
			assert.Contains(t, err.Error(), tc.expectedErr.Error(), fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		} else {
			assert.Nil(t, err, fmt.Sprintf("%s: expected error: %v, got: %v", tc.name, tc.expectedErr, err))
		}
		assert.Equal(t, tc.expectedResp, resp, fmt.Sprintf("expected response: %v, got: %v", tc.expectedResp, resp))

		call.Unset()
	}
}
