package http

import (
	"fmt"

	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// createRoom creates a new chat room.
// @Summary Create a new chat room
// @Description Allows authenticated users to create a new chat room.
// @Tags Room
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param   body body dto.CreateRoomReqDTO true "Create Room"
// @Success 200 {object} values.RoomValue
// @Failure 400 {object} errorhandler.Response
// @Router /chat/room [post]
func (h *handler) createRoom() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request dto.CreateRoomReqDTO
		err := f.BodyParser(&request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.RequestBodyParseErrorMessage,
				zap.Error(err),
			)

			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.RequestBodyParseErrorCode, Message: errorhandler.RequestBodyParseErrorMessage, StatusCode: fiber.StatusBadRequest})
		}

		request.Owner = f.Locals("userID").(string)

		validate := validator.New()

		// Validate the request
		err = validate.Struct(request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.ValidationErrorMessage,
				zap.Error(err),
			)

			fields := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				fields = append(fields, err.Field())
			}
			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.ValidationErrorCode, Message: fmt.Sprintf("invalid fields %s", fields), StatusCode: fiber.StatusBadRequest})
		}

		// Handle the request
		res, resErr := h.application.RoomCommandHandler.CreateRoom(&request)

		if resErr != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.InvalidCredentialsMessage,
				zap.Error(err),
			)

			return f.Status(resErr.StatusCode).JSON(resErr)
		}

		return f.Status(fiber.StatusOK).JSON(res)
	}
}

// queryRooms gets a list of chat rooms based on the provided criteria.
// @Summary Query chat rooms
// @Description Retrieves a list of chat rooms based on filters, sorting and pagination.
// @Tags Room
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param   page query int false "Page number for pagination" default(1)
// @Param   per_page query int false "Number of items per page" default(10)
// @Param   sort_field query string false "Field to sort by" default("created_at")
// @Param   sort_order query string false "Order of sorting" Enums(asc, desc) default("desc")
// @Param   user_in query string false "Filter rooms where user is in"
// @Param   user_not_in query string false "Filter rooms where user is not in"
// @Success 200 {array} []values.RoomValue
// @Failure 400 {object} errorhandler.Response
// @Router /chat/rooms [get]
func (h *handler) queryRooms() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request = &dto.QueryRoomsReqDTO{
			Page:      f.QueryInt("page", 1),
			PerPage:   f.QueryInt("per_page", 10),
			SortField: f.Query("sort_field", "created_at"),
			SortOrder: f.Query("sort_order", "desc"),
		}

		// Filters
		if f.Query("user_in") != "" {
			request.UserIn = []string{f.Query("user_in")}

			// Add the current user to the filter
			request.UserIn = append(request.UserIn, f.Locals("userID").(string))
		}

		if f.Query("user_not_in") != "" {
			request.UserNotIn = []string{f.Query("user_not_in")}
		}

		validate := validator.New()

		// Validate the request
		err := validate.Struct(request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.ValidationErrorMessage,
				zap.Error(err),
			)

			fields := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				fields = append(fields, err.Field())
			}
			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.ValidationErrorCode, Message: fmt.Sprintf("invalid fields %s", fields), StatusCode: fiber.StatusBadRequest})
		}

		// Handle the request
		res, resErr := h.application.RoomQueryHandler.QueryRooms(request)

		if resErr != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.InvalidCredentialsMessage,
				zap.Error(err),
			)

			return f.Status(resErr.StatusCode).JSON(resErr)
		}

		return f.Status(fiber.StatusOK).JSON(res)
	}
}

// joinRoom joins a user to a chat room.
// @Summary Join a chat room
// @Description Allows authenticated users to join a chat room.
// @Tags Room
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param   roomID path string true "Room ID"
// @Success 200 {object} values.RoomValue
// @Failure 400 {object} errorhandler.Response
// @Router /chat/room/{roomID}/join [post]
func (h *handler) joinRoom() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request = &dto.JoinRoomReqDTO{
			RoomID: f.Params("roomID"),
			UserID: f.Locals("userID").(string),
		}

		validate := validator.New()

		// Validate the request
		err := validate.Struct(request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.ValidationErrorMessage,
				zap.Error(err),
			)

			fields := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				fields = append(fields, err.Field())
			}
			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.ValidationErrorCode, Message: fmt.Sprintf("invalid fields %s", fields), StatusCode: fiber.StatusBadRequest})
		}

		// Handle the request
		room, resErr := h.application.RoomCommandHandler.JoinRoom(request)

		if resErr != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.InvalidCredentialsMessage,
				zap.Error(err),
			)

			return f.Status(resErr.StatusCode).JSON(resErr)
		}

		return f.Status(fiber.StatusOK).JSON(room)
	}
}

// leaveRoom removes a user from a chat room.
// @Summary Leave a chat room
// @Description Allows authenticated users to leave a chat room.
// @Tags Room
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param   roomID path string true "Room ID"
// @Success 200 {object} values.RoomValue
// @Failure 400 {object} errorhandler.Response
// @Router /chat/room/{roomID}/leave [post]
func (h *handler) leaveRoom() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request = &dto.LeaveRoomReqDTO{
			RoomID: f.Params("roomID"),
			UserID: f.Locals("userID").(string),
		}

		validate := validator.New()

		// Validate the request
		err := validate.Struct(request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.ValidationErrorMessage,
				zap.Error(err),
			)

			fields := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				fields = append(fields, err.Field())
			}
			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.ValidationErrorCode, Message: fmt.Sprintf("invalid fields %s", fields), StatusCode: fiber.StatusBadRequest})
		}

		// Handle the request
		_, resErr := h.application.RoomCommandHandler.LeaveRoom(request)

		if resErr != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.InvalidCredentialsMessage,
				zap.Error(err),
			)

			return f.Status(resErr.StatusCode).JSON(resErr)
		}

		return f.Status(fiber.StatusOK).JSON(nil)
	}
}
