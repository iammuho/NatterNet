package http

import (
	"fmt"

	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// createMessage  is the handler for creating a message
func (h *handler) createMessage() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request dto.CreateMessageReqDTO
		err := f.BodyParser(&request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.RequestBodyParseErrorMessage,
				zap.Error(err),
			)

			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.RequestBodyParseErrorCode, Message: errorhandler.RequestBodyParseErrorMessage, StatusCode: fiber.StatusBadRequest})
		}

		request.RoomID = f.Params("roomID")
		request.SenderID = f.Locals("userID").(string)

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
		res, resErr := h.application.MessageCommandHandler.CreateMessage(&request)

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

// queryMessages is the handler for querying messages
func (h *handler) queryMessages() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		request := &dto.QueryMessagesReqDTO{
			RoomID:    f.Params("roomID"),
			UserID:    f.Locals("userID").(string),
			Page:      f.QueryInt("page", 1),
			PerPage:   f.QueryInt("per_page", 10),
			SortField: f.Query("sort_field", "created_at"),
			SortOrder: f.Query("sort_order", "desc"),
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
		res, resErr := h.application.MessageQueryHandler.QueryMessages(request)

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
