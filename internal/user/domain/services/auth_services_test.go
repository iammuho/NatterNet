package services

import (
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	mockcontext "github.com/iammuho/natternet/cmd/app/context/mocks"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/internal/user/domain/entity"
	mockuserrepository "github.com/iammuho/natternet/internal/user/domain/repository/mocks"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
	mockhashingfactory "github.com/iammuho/natternet/pkg/hashing/mocks"
	mockutils "github.com/iammuho/natternet/pkg/utils/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthDomainServices_SignIn(t *testing.T) {
	justNow := time.Now()
	uuid := "some-uuid"

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
			mockUser:          entity.NewUser(uuid, "testuser", "hashed_password", "test@test.com", justNow),
			expectedErr:       false,
			expectedUserValue: &values.UserValue{ID: uuid, Username: "testuser", Email: "test@test.com", CreatedAt: justNow},
			wrongPassword:     false,
		},
		{
			name:              "Invalid Credentials",
			input:             &dto.SignInReqDTO{Login: "testuser", Password: "wrong_password"},
			mockUser:          entity.NewUser(uuid, "testuser", "hashed_password", "test@test.com", justNow),
			expectedUserValue: &values.UserValue{ID: uuid, Username: "testuser", Email: "test@test.com"},
			expectedErr:       true,
			mockUserErr:       &errorhandler.Response{Code: errorhandler.InvalidCredentialsErrorCode, Message: errorhandler.InvalidCredentialsMessage},
			wrongPassword:     true,
		},
		{
			name:              "User Not Found",
			input:             &dto.SignInReqDTO{Login: "testuser", Password: "hashed_password"},
			mockUser:          entity.NewUser(uuid, "testuser", "hashed_password", "test@test.com", justNow),
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
			mockHashingFactory.EXPECT().ComparePassword(test.input.Password, gomock.Any()).Return(!test.wrongPassword).AnyTimes()
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

func TestAuthDomainServices_SignUp(t *testing.T) {
	justNow := time.Now()
	uuid := "some-uuid"

	tests := []struct {
		name               string
		input              *dto.SignupReqDTO
		mockUserByEmail    *values.UserDBValue
		mockUserByUsername *values.UserDBValue
		expectedErr        *errorhandler.Response
		expectedUser       *values.UserValue
	}{
		{
			name:               "Successful Sign Up",
			input:              &dto.SignupReqDTO{Username: "testuser", Email: "test@test.com", Password: "hashed_password"},
			mockUserByEmail:    nil,
			mockUserByUsername: nil,
			expectedUser:       &values.UserValue{ID: uuid, Username: "testuser", Email: "test@test.com", CreatedAt: justNow},
		},
		{
			name:            "Email Already Exists",
			input:           &dto.SignupReqDTO{Username: "testuser", Email: "test@test.com", Password: "hashed_password"},
			mockUserByEmail: values.NewUserDBValueFromUser(entity.NewUser(uuid, "testuser", "hashed_password", "test@test.com", justNow)),
			expectedErr:     &errorhandler.Response{Code: errorhandler.EmailAlreadyExistsErrorCode, Message: errorhandler.EmailAlreadyExistsMessage, StatusCode: fiber.StatusConflict},
		},
		{
			name:               "Username Already Exists",
			input:              &dto.SignupReqDTO{Username: "testuser", Email: "test@test.com", Password: "hashed_password"},
			mockUserByUsername: values.NewUserDBValueFromUser(entity.NewUser(uuid, "testuser", "hashed_password", "test@test.com", justNow)),
			expectedErr:        &errorhandler.Response{Code: errorhandler.UsernameAlreadyExistsErrorCode, Message: errorhandler.UsernameAlreadyExistsMessage, StatusCode: fiber.StatusConflict},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAppContext := mockcontext.NewMockAppContext(ctrl)
			mockUserRepository := mockuserrepository.NewMockUserRepository(ctrl)

			authService := NewAuthDomainServices(mockAppContext, mockUserRepository)

			mockUUID := mockutils.NewMockUUID(ctrl)
			mockUUID.EXPECT().NewUUID().Return(uuid).AnyTimes()
			mockAppContext.EXPECT().GetUUID().Return(mockUUID).AnyTimes()

			mockTimer := mockutils.NewMockTimer(ctrl)
			mockTimer.EXPECT().Now().Return(justNow).AnyTimes()
			mockAppContext.EXPECT().GetTimer().Return(mockTimer).AnyTimes()

			mockUserRepository.EXPECT().FindOneByEmail(test.input.Email).Return(test.mockUserByEmail, test.expectedErr).AnyTimes()
			mockUserRepository.EXPECT().FindOneByUsername(test.input.Username).Return(test.mockUserByUsername, test.expectedErr).AnyTimes()

			if test.expectedErr == nil {
				mockUserRepository.EXPECT().Create(gomock.Any()).Return(nil).Times(1)
			}

			resp, err := authService.SignUp(test.input)

			if test.expectedErr != nil {
				assert.Nil(t, resp)
				assert.NotNil(t, err)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, test.expectedUser, resp)
			}
		})
	}
}
