package repository

import (
	"github.com/iammuho/natternet/internal/chat/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type RoomRepository interface {
	// Commands
	Create(room *values.RoomDBValue) *errorhandler.Response
}
