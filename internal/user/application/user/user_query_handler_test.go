package user

import (
	"testing"
	"time"

	mockcontext "github.com/iammuho/natternet/cmd/app/context/mocks"
	"github.com/iammuho/natternet/internal/user/application/user/dto"
	mockuserdomainservices "github.com/iammuho/natternet/internal/user/domain/services/mocks"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserQueryHandler_QueryUserByID(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name              string
		input             *dto.QueryUserByIDReqDTO
		mockUserResp      *values.UserValue
		mockUserErr       *errorhandler.Response
		expectedErr       bool
		expectedUserValue *values.UserValue
	}{
		{
			name:              "Successful Query",
			input:             &dto.QueryUserByIDReqDTO{UserID: "1"},
			mockUserResp:      &values.UserValue{ID: "1", Username: "testuser", Email: "test@example.com", CreatedAt: now},
			expectedErr:       false,
			expectedUserValue: &values.UserValue{ID: "1", Username: "testuser", Email: "test@example.com", CreatedAt: now},
		},
		{
			name:        "User Not Found Error",
			input:       &dto.QueryUserByIDReqDTO{UserID: "2"},
			mockUserErr: &errorhandler.Response{Code: 404, Message: "User not found"},
			expectedErr: true,
		},
		{
			name:        "Invalid UserID Error",
			input:       &dto.QueryUserByIDReqDTO{UserID: ""},
			mockUserErr: &errorhandler.Response{Code: 400, Message: "Invalid UserID"},
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAppContext := mockcontext.NewMockAppContext(ctrl)
			mockUserQueryDomainServices := mockuserdomainservices.NewMockUserQueryDomainServices(ctrl)

			mockUserQueryDomainServices.EXPECT().FindByID(test.input.UserID).Return(test.mockUserResp, test.mockUserErr).Times(1)

			queryHandler := NewUserQueryHandler(mockAppContext, mockUserQueryDomainServices)
			resp, err := queryHandler.QueryUserByID(test.input)

			if test.expectedErr {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				} else if err.Message != test.mockUserErr.Message {
					t.Errorf("Expected error message to be: %s, but got: %s", test.mockUserErr.Message, err.Message)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
				assert.Equal(t, test.expectedUserValue, resp)
			}
		})
	}
}
