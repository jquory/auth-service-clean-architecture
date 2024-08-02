package middleware

import (
	"auth-service/app/common/logs"
	"auth-service/app/dto"
	"auth-service/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateJWT(role string, cfg config.Config) func(ctx *fiber.Ctx) error {
	jwtSecret := cfg.Jwt.JwtSecret
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: jwtSecret},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			roles := claims["roles"].([]interface{})

			logs.NewLogger().Info("Role user", role)
			for _, roleInterface := range roles {
				roleMap := roleInterface.(map[string]interface{})

				if roleMap["role"] == role {
					return ctx.Next()
				}
			}

			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.GeneralResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       "Invalid Role",
			})
		},

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return ctx.Status(fiber.StatusBadRequest).JSON(dto.GeneralResponse{
					StatusCode: fiber.StatusBadRequest,
					Message:    err.Error(),
				})
			}
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.GeneralResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       "Token invalid or expired",
			})
		},
	})
}
