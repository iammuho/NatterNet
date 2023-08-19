package http

import (
	"fmt"

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
		var request dto.SignInReqDTO
		err := f.BodyParser(&request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.RequestBodyParseErrorMessage,
				zap.Error(err),
			)

			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.RequestBodyParseErrorCode, Message: errorhandler.RequestBodyParseErrorMessage, StatusCode: fiber.StatusBadRequest})
		}

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
		res, resErr := h.application.SigninCommandHandler.Handle(&request)

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

// Signup is the handler for the signup route
func (h *handler) Signup() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		var request dto.SignupReqDTO
		err := f.BodyParser(&request)
		if err != nil {
			h.application.AppContext.GetLogger().Logger.Warn(
				errorhandler.RequestBodyParseErrorMessage,
				zap.Error(err),
			)

			return f.Status(fiber.StatusBadRequest).JSON(&errorhandler.Response{Code: errorhandler.RequestBodyParseErrorCode, Message: errorhandler.RequestBodyParseErrorMessage, StatusCode: fiber.StatusBadRequest})
		}

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
		res, resErr := h.application.SignupCommandHandler.Handle(&request)

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
