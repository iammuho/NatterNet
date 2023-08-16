package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

// Middleware is the struct pool for the user domain middleware
type Middleware struct {
	ctx context.AppContext
}

// NewAuthMiddleware is the constructor for the user domain middleware
func NewMiddleware(ctx context.AppContext) *Middleware {
	return &Middleware{
		ctx: ctx,
	}
}

// Protected protect routes
func (h *Middleware) Protected() fiber.Handler {
	return func(f *fiber.Ctx) error {
		if f.Request().Header.Peek("Authorization") == nil {
			return f.Status(fiber.StatusUnauthorized).
				JSON(&errorhandler.Response{Code: errorhandler.InvalidAccessTokenErrorCode, Message: errorhandler.InvalidAccessTokenMessage, StatusCode: fiber.StatusUnauthorized})
		}

		// Parse the JWT
		claims, err := h.ctx.GetJwtContext().ParseJWT(string(f.Request().Header.Peek("Authorization")))

		// Find active user
		if err != nil {
			return f.Status(err.StatusCode).JSON(err)
		}

		if claims == nil || claims["ID"] == nil {
			return f.Status(fiber.StatusUnauthorized).
				JSON(&errorhandler.Response{Code: errorhandler.InvalidAccessTokenErrorCode, Message: errorhandler.InvalidAccessTokenMessage, StatusCode: fiber.StatusUnauthorized})
		}

		f.Locals("userID", claims["ID"].(string))
		f.Next()

		return nil
	}
}
