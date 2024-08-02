package exceptions

import (
	"e-rt/app/dto"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var validationError ValidationError
	valErr := errors.As(err, &validationError)
	if valErr {
		data := err.Error()
		var message []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &message)
		PanicLogging(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Data:       message,
		})
	}

	var notFoundError NotFoundError
	notFound := errors.As(err, &notFoundError)
	if notFound {
		data := err.Error()
		var message []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &message)
		PanicLogging(errJson)
		return ctx.Status(fiber.StatusNotFound).JSON(dto.GeneralResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
			Data:       message,
		})
	}

	var unAuthorizedError UnAuthorizedError
	unAuthorized := errors.As(err, &unAuthorizedError)
	if unAuthorized {
		data := err.Error()
		var message []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &message)
		PanicLogging(errJson)
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.GeneralResponse{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
			Data:       message,
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(dto.GeneralResponse{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "Internal Server Error",
		Data:       err.Error(),
	})
}
