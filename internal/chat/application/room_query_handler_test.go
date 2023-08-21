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

func TestRoomQueryHandler_CreateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAppContext := mockcontext.NewMockAppContext(ctrl)
	mockRoomQueryServices := mockchatdomainservices.NewMockRoomQueryDomainServices(ctrl)

	roomQueryHandler := NewRoomQueryHandler(mockAppContext, mockRoomQueryServices)

	tests := []struct {
		name           string
		req            *dto.QueryRoomsReqDTO
		mockRoomsValue []*values.RoomValue
		mockErr        *errorhandler.Response
		expectedRooms  []*values.RoomValue
		expectedErr    *errorhandler.Response
	}{
		{
			name: "Successful Room Query",
			req:  &dto.QueryRoomsReqDTO{UserIn: []string{"1", "2"}},
			mockRoomsValue: []*values.RoomValue{
				{
					ID: "1",
					Meta: entity.RoomMeta{
						Name:        "Test Room",
						Description: "This is a test room",
					},
				},
			},
			expectedRooms: []*values.RoomValue{
				{
					ID: "1",
					Meta: entity.RoomMeta{
						Name:        "Test Room",
						Description: "This is a test room",
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:           "Failed Room Query",
			req:            &dto.QueryRoomsReqDTO{UserIn: []string{"1", "2"}},
			mockRoomsValue: nil,
			mockErr:        &errorhandler.Response{Code: errorhandler.InternalSystemErrorCode, Message: "Internal Server Error"},
			expectedRooms:  nil,
			expectedErr:    &errorhandler.Response{Code: errorhandler.InternalSystemErrorCode, Message: "Internal Server Error"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRoomQueryServices.EXPECT().QueryRooms(test.req).Return(test.mockRoomsValue, test.mockErr).Times(1)

			room, err := roomQueryHandler.QueryRooms(test.req)

			assert.Equal(t, test.expectedRooms, room)

			if test.expectedErr != nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
