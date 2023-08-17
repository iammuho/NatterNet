package auth

import (
	"testing"

	mockcontext "github.com/iammuho/natternet/cmd/app/context/mocks"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	mockauthdomainservices "github.com/iammuho/natternet/internal/user/domain/services/mocks"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
	"github.com/iammuho/natternet/pkg/jwt"
	mockjwt "github.com/iammuho/natternet/pkg/jwt/mocks"

	"go.uber.org/mock/gomock"
)

func TestSignUpCommandHandler_Handle(t *testing.T) {
	tests := []struct {
		name           string
		input          *dto.SignupReqDTO
		mockSignUpResp *values.UserValue
		mockSignUpErr  *errorhandler.Response
		mockJwtResp    *jwt.JWTResponse
		mockJwtErr     *errorhandler.Response
		expectedErr    bool
		expectedToken  string
		expectedExpire int
	}{
		{
			name:           "Successful Signup",
			input:          &dto.SignupReqDTO{Email: "test@example.com", Username: "testuser", Password: "12345678"},
			mockSignUpResp: &values.UserValue{ID: "1"},
			mockJwtResp:    &jwt.JWTResponse{AccessToken: "test-access-token", RefreshToken: "test-refresh-token", ExpiresIn: 123456},
			expectedErr:    false,
			expectedToken:  "test-access-token",
			expectedExpire: 123456,
		},
		{
			name:           "SignUp Error",
			input:          &dto.SignupReqDTO{Email: "test@example.com", Username: "testuser", Password: "12345678"},
			mockSignUpResp: &values.UserValue{ID: "1"},
			mockSignUpErr:  &errorhandler.Response{Code: 500, Message: "some error"},
			expectedErr:    true,
		},
		{
			name:           "CreatePair Error",
			input:          &dto.SignupReqDTO{Email: "test@example.com", Username: "testuser", Password: "12345678"},
			mockSignUpResp: &values.UserValue{ID: "1"},
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
			mockAuthService := mockauthdomainservices.NewMockAuthDomainServices(ctrl)

			mockAppContext.EXPECT().GetJwtContext().Return(mockJwtContext).AnyTimes()
			mockAuthService.EXPECT().SignUp(test.input).Return(test.mockSignUpResp, test.mockSignUpErr).Times(1)

			if test.mockSignUpErr == nil {
				mockJwtContext.EXPECT().CreatePair(gomock.Any()).Return(test.mockJwtResp, test.mockJwtErr).Times(1)
			}

			cmd := NewSignUpCommandHandler(mockAppContext, mockAuthService)
			resp, err := cmd.Handle(test.input)

			if test.expectedErr {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				} else if test.mockJwtErr != nil && err.Code != test.mockJwtErr.Code {
					t.Errorf("Expected error message to be: %s, but got: %s", test.mockJwtErr.Message, err.Message)
				} else if test.mockSignUpErr != nil && err.Code != test.mockSignUpErr.Code {
					t.Errorf("Expected error message to be: %s, but got: %s", test.mockSignUpErr.Message, err.Message)
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
