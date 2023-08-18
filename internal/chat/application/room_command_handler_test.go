package application

import (
	"testing"

	mockcontext "github.com/iammuho/natternet/cmd/app/context/mocks"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/entity"
	mockchatdomainservices "github.com/iammuho/natternet/internal/chat/domain/services/mocks"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRoomCommandHandler_CreateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAppContext := mockcontext.NewMockAppContext(ctrl)
	mockRoomCommandServices := mockchatdomainservices.NewMockRoomCommandDomainServices(ctrl)

	roomCommandHandler := NewRoomCommandHandler(mockAppContext, mockRoomCommandServices)

	tests := []struct {
		name          string
		req           *dto.CreateRoomReqDTO
		mockRoomValue *values.RoomValue
		mockErr       *errorhandler.Response
		expectedRoom  *values.RoomValue
		expectedErr   *errorhandler.Response
	}{
		{
			name: "Successful Room Creation",
			req:  &dto.CreateRoomReqDTO{Name: "Test Room", Description: "This is a test room", IsGroup: true, UserIDs: []string{"1", "2"}},
			mockRoomValue: &values.RoomValue{
				ID: "1",
				Meta: entity.RoomMeta{
					Name:        "Test Room",
					Description: "This is a test room",
				},
			},
			expectedRoom: &values.RoomValue{
				ID: "1",
				Meta: entity.RoomMeta{
					Name:        "Test Room",
					Description: "This is a test room",
				},
			},
			expectedErr: nil,
		},
		{
			name:          "Failed Room Creation",
			req:           &dto.CreateRoomReqDTO{Name: "Test Room", Description: "This is a test room", IsGroup: true, UserIDs: []string{"1", "2"}},
			mockRoomValue: nil,
			mockErr:       &errorhandler.Response{Code: errorhandler.InternalSystemErrorCode, Message: "Internal Server Error"},
			expectedRoom:  nil,
			expectedErr:   &errorhandler.Response{Code: errorhandler.InternalSystemErrorCode, Message: "Internal Server Error"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRoomCommandServices.EXPECT().CreateRoom(test.req).Return(test.mockRoomValue, test.mockErr).Times(1)

			room, err := roomCommandHandler.CreateRoom(test.req)

			assert.Equal(t, test.expectedRoom, room)

			if test.expectedErr != nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
