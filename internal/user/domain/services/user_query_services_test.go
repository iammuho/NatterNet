package services

import (
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	mockcontext "github.com/iammuho/natternet/cmd/app/context/mocks"
	"github.com/iammuho/natternet/internal/user/domain/entity"
	mockuserrepository "github.com/iammuho/natternet/internal/user/domain/repository/mocks"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserQueryDomainServices_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAppContext := mockcontext.NewMockAppContext(ctrl)
	mockUserRepository := mockuserrepository.NewMockUserRepository(ctrl)

	userQueryService := NewUserQueryDomainServices(mockAppContext, mockUserRepository)

	justNow := time.Now()
	mockUser := entity.NewUser("1", "testuser", "hashed_password", "test@test.com", justNow)
	mockUserValue := values.NewUserValueFromUser(mockUser)

	tests := []struct {
		name          string
		id            string
		mockUser      *values.UserDBValue
		mockErr       *errorhandler.Response
		expectedUser  *values.UserValue
		expectedError *errorhandler.Response
	}{
		{
			name:          "User Found",
			id:            "1",
			mockUser:      values.NewUserDBValueFromUser(mockUser),
			expectedUser:  mockUserValue,
			expectedError: nil,
		},
		{
			name:          "User Not Found",
			id:            "2",
			mockUser:      nil,
			expectedUser:  nil,
			expectedError: &errorhandler.Response{Code: errorhandler.UserNotFoundErrorCode, Message: errorhandler.UserNotFoundMessage, StatusCode: fiber.StatusNotFound},
		},
		{
			name:          "Error Finding User",
			id:            "1",
			mockUser:      nil,
			mockErr:       &errorhandler.Response{Code: errorhandler.InternalSystemErrorCode, Message: "Internal Server Error"},
			expectedUser:  nil,
			expectedError: &errorhandler.Response{Code: errorhandler.InternalSystemErrorCode, Message: "Internal Server Error"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockUserRepository.EXPECT().FindOneByID(test.id).Return(test.mockUser, test.mockErr).Times(1)

			user, err := userQueryService.FindByID(test.id)

			assert.Equal(t, test.expectedUser, user)

			if test.expectedError != nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
