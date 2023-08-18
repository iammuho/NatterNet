package services

import (
	"testing"
	"time"

	mockcontext "github.com/iammuho/natternet/cmd/app/context/mocks"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/entity"
	mockuserrepository "github.com/iammuho/natternet/internal/user/domain/repository/mocks"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
	mockhashingfactory "github.com/iammuho/natternet/pkg/hashing/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthDomainServices_SignIn(t *testing.T) {
	justNow := time.Now()

	tests := []struct {
		name              string
		input             *dto.SignInReqDTO
		mockUser          *entity.User
		mockUserErr       *errorhandler.Response
		expectedErr       bool
		expectedUserValue *values.UserValue
		wrongPassword     bool
	}{
		{
			name:              "Successful Sign In",
			input:             &dto.SignInReqDTO{Login: "testuser", Password: "hashed_password"},
			mockUser:          entity.NewUser("1", "testuser", "hashed_password", "test@test.com", justNow),
			expectedErr:       false,
			expectedUserValue: &values.UserValue{ID: "1", Username: "testuser", Email: "test@test.com", CreatedAt: justNow},
			wrongPassword:     false,
		},
		{
			name:              "Invalid Credentials",
			input:             &dto.SignInReqDTO{Login: "testuser", Password: "wrong_password"},
			mockUser:          entity.NewUser("1", "testuser", "hashed_password", "test@test.com", justNow),
			expectedUserValue: &values.UserValue{ID: "1", Username: "testuser", Email: "test@test.com"},
			expectedErr:       true,
			mockUserErr:       &errorhandler.Response{Code: errorhandler.InvalidCredentialsErrorCode, Message: errorhandler.InvalidCredentialsMessage},
			wrongPassword:     true,
		},
		{
			name:              "User Not Found",
			input:             &dto.SignInReqDTO{Login: "testuser", Password: "hashed_password"},
			mockUser:          entity.NewUser("1", "testuser", "hashed_password", "test@test.com", justNow),
			mockUserErr:       &errorhandler.Response{Code: errorhandler.UserNotFoundErrorCode, Message: errorhandler.UserNotFoundMessage},
			expectedUserValue: nil,
			expectedErr:       true,
			wrongPassword:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAppContext := mockcontext.NewMockAppContext(ctrl)
			mockUserRepository := mockuserrepository.NewMockUserRepository(ctrl)

			authService := NewAuthDomainServices(mockAppContext, mockUserRepository)

			// Mock the mockcontext.MockAppContext.GetHashingFactory
			mockHashingFactory := mockhashingfactory.NewMockHashingFactory(ctrl)
			mockHashingFactory.EXPECT().ComparePassword(gomock.Any(), test.input.Password).Return(!test.wrongPassword).AnyTimes()
			mockAppContext.EXPECT().GetHashingFactory().Return(mockHashingFactory).AnyTimes()

			mockUserRepository.EXPECT().FindOneByLogin(test.input.Login).Return(values.NewUserDBValueFromUser(test.mockUser), test.mockUserErr).Times(1)

			resp, err := authService.SignIn(test.input)

			if test.expectedErr {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}

				assert.Equal(t, test.mockUserErr, err)
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}

				assert.Equal(t, test.expectedUserValue, resp)
			}
		})
	}
}
