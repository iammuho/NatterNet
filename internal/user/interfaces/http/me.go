package http

import (
	"fmt"

	"github.com/iammuho/natternet/internal/user/application/user"
	userDTO "github.com/iammuho/natternet/internal/user/application/user/dto"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/internal/user/infrastructure/mongodb"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Me is the handler for the me route to get the user's own information
func (h *handler) Me() fiber.Handler {
	return func(f *fiber.Ctx) error {
		// Serialize the body
		request := &userDTO.QueryUserByIDReqDTO{
			UserID: f.Locals("userID").(string),
		}

		validate := validator.New()

		// Validate the request
		err := validate.Struct(request)
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

		// initialize the user repository
		userRepository := mongodb.NewUserRepository(h.ctx)

		// Initialize the userQueryDomainServices
		userQueryDomainServices := services.NewUserQueryDomainServices(h.ctx, userRepository)

		// Setup the command handlers
		userQueryHandler := user.NewUserQueryHandler(h.ctx, userQueryDomainServices)

		// Handle the request
		res, resErr := userQueryHandler.QueryUserByID(request)

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
