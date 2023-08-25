package http

import (
	"fmt"

	userDTO "github.com/iammuho/natternet/internal/user/application/user/dto"
	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Me returns the authenticated user's details.
// @Summary User's own details
// @Description Get the authenticated user's own details.
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} values.UserValue
// @Failure 400 {object} errorhandler.Response
// @Router /user/me [get]
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
		res, resErr := h.application.UserQueryHandler.QueryUserByID(request)

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
