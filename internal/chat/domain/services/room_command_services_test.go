package services

import (
	"testing"
	"time"

	mockcontext "github.com/iammuho/natternet/cmd/app/context/mocks"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/internal/chat/domain/entity"
	mockchatrepository "github.com/iammuho/natternet/internal/chat/domain/repository/mocks"
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
	mockutils "github.com/iammuho/natternet/pkg/utils/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRoomCommandDomainServices_CreateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAppContext := mockcontext.NewMockAppContext(ctrl)
	mockUUID := mockutils.NewMockUUID(ctrl)
	mockTimer := mockutils.NewMockTimer(ctrl)
	mockRoomRepository := mockchatrepository.NewMockRoomRepository(ctrl)

	roomCommandDomainService := NewRoomCommandDomainServices(mockAppContext, mockRoomRepository)

	tests := []struct {
		name          string
		req           *dto.CreateRoomReqDTO
		expectedRoom  *values.RoomValue
		expectedError *errorhandler.Response
	}{
		{
			name: "Successful Room Creation",
			req: &dto.CreateRoomReqDTO{
				Name:        "Test Room",
				Description: "This is a test room",
				IsGroup:     true,
				UserIDs:     []string{"1", "2"},
			},
			expectedRoom: &values.RoomValue{
				ID: "1",
				Meta: entity.RoomMeta{
					Name:        "Test Room",
					Description: "This is a test room",
				},
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Generate a fixed UUID and time
			uuid := "roomID"
			createdAt := time.Now()

			// Set up mock expectations
			mockAppContext.EXPECT().GetUUID().Return(mockUUID).Times(1)
			mockUUID.EXPECT().NewUUID().Return(uuid).Times(1)

			mockAppContext.EXPECT().GetTimer().Return(mockTimer).Times(1)
			mockTimer.EXPECT().Now().Return(createdAt).Times(1)

			if test.expectedError != nil {
				// Mock the roomRepository to simulate an error
				mockRoomRepository.EXPECT().Create(gomock.Any()).Return(test.expectedError).Times(1)
			} else {
				// Mock the roomRepository to simulate a successful creation
				mockRoomRepository.EXPECT().Create(gomock.Any()).Return(nil).Times(1)
			}

			// Call the CreateRoom function
			room, err := roomCommandDomainService.CreateRoom(test.req)

			// Assertions
			assert.Equal(t, test.expectedRoom, room)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
