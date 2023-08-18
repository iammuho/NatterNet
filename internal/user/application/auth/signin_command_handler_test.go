package auth

import (
	"testing"

	mockcontext "github.com/iammuho/natternet/cmd/app/context/mocks"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	mockuserdomainservices "github.com/iammuho/natternet/internal/user/domain/services/mocks"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"github.com/iammuho/natternet/pkg/jwt"
	mockjwt "github.com/iammuho/natternet/pkg/jwt/mocks"

	"go.uber.org/mock/gomock"
)

func TestSignInCommandHandler_Handle(t *testing.T) {
	tests := []struct {
		name           string
		input          *dto.SignInReqDTO
		mockSignInResp *values.UserValue
		mockSignInErr  *errorhandler.Response
		mockJwtResp    *jwt.JWTResponse
		mockJwtErr     *errorhandler.Response
		expectedErr    bool
		expectedToken  string
		expectedExpire int
	}{
		{
			name:           "Successful Signin",
			input:          &dto.SignInReqDTO{Login: "test", Password: "123456"},
			mockSignInResp: &values.UserValue{ID: "1"},
			mockJwtResp:    &jwt.JWTResponse{AccessToken: "test-access-token", RefreshToken: "test-refresh-token", ExpiresIn: 123456},
			expectedErr:    false,
			expectedToken:  "test-access-token",
			expectedExpire: 123456,
		},
		{
			name:          "SignIn Error",
			input:         &dto.SignInReqDTO{Login: "test", Password: "123456"},
			mockSignInErr: &errorhandler.Response{Code: 500, Message: "some error"},
			expectedErr:   true,
		},
		{
			name:           "CreatePair Error",
			input:          &dto.SignInReqDTO{Login: "test", Password: "123456"},
			mockSignInResp: &values.UserValue{ID: "1"},
			mockJwtErr:     &errorhandler.Response{Code: 500, Message: "some error"},
			expectedErr:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAppContext := mockcontext.NewMockAppContext(ctrl)
			mockJwtContext := mockjwt.NewMockJwtContext(ctrl)
			mockAuthService := mockuserdomainservices.NewMockAuthDomainServices(ctrl)
			mockAppContext.EXPECT().GetJwtContext().Return(mockJwtContext).AnyTimes()
			mockAuthService.EXPECT().SignIn(test.input).Return(test.mockSignInResp, test.mockSignInErr).Times(1)

			if test.mockSignInErr == nil {
				mockJwtContext.EXPECT().CreatePair(gomock.Any()).Return(test.mockJwtResp, test.mockJwtErr).Times(1)
			}

			cmd := NewSignInCommandHandler(mockAppContext, mockAuthService)
			resp, err := cmd.Handle(test.input)

			if test.expectedErr {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}

				if resp.AccessToken != test.expectedToken {
					t.Errorf("Expected access token to be: %s, but got: %s", test.expectedToken, resp.AccessToken)
				}
			}
		})
	}
}
