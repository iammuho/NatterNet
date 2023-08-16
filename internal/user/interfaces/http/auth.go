package http

import (
	"fmt"

	"github.com/iammuho/natternet/internal/user/application/auth"
	"github.com/iammuho/natternet/internal/user/application/auth/dto"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Signin is the handler for the signin route
func (h *handler) Signin() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request dto.SigninReqDTO
		err := f.BodyParser(&request)
		if err != nil {
			h.ctx.GetLogger().Logger.Warn(
				errorhandler.RequestBodyParseErrorMessage,
				zap.Error(err),
			)

			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.RequestBodyParseErrorCode, Message: errorhandler.RequestBodyParseErrorMessage, StatusCode: fiber.StatusBadRequest})
		}

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
		signinCommandHandler := auth.NewSigninCommandHandler(h.ctx)

		// Handle the request
		res, resErr := signinCommandHandler.Handle(&request)

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
