package http

import (
	"fmt"

	"github.com/iammuho/natternet/internal/chat/application"
	"github.com/iammuho/natternet/internal/chat/application/dto"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// CreateRoom is the handler for the create room path
func (h *handler) CreateRoom() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request dto.CreateRoomReqDTO
		err := f.BodyParser(&request)
		if err != nil {
			h.ctx.GetLogger().Logger.Warn(
				errorhandler.RequestBodyParseErrorMessage,
				zap.Error(err),
			)

			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.RequestBodyParseErrorCode, Message: errorhandler.RequestBodyParseErrorMessage, StatusCode: fiber.StatusBadRequest})
		}

		request.UserIDs = append(request.UserIDs, f.Locals("userID").(string))

		validate := validator.New()

		// Validate the request
		err = validate.Struct(request)
		if err != nil {
			h.ctx.GetLogger().Logger.Warn(
				errorhandler.ValidationErrorMessage,
				zap.Error(err),
			)

			fields := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				fields = append(fields, err.Field())
			}
			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.ValidationErrorCode, Message: fmt.Sprintf("invalid fields %s", fields), StatusCode: fiber.StatusBadRequest})
		}

		// Setup the command handlers
		roomCommandHandler := application.NewRoomCommandHandler(h.ctx)

		// Handle the request
		res, resErr := roomCommandHandler.CreateRoom(&request)

		if resErr != nil {
			h.ctx.GetLogger().Logger.Warn(
				errorhandler.InvalidCredentialsMessage,
				zap.Error(err),
			)

			return f.Status(resErr.StatusCode).JSON(resErr)
		}

		return f.Status(fiber.StatusOK).JSON(res)
	}
}
