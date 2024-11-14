package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type APIError struct {
	Code int `json:"code"`
	Err  any `json:"err"`
}

func (e APIError) Error() string {
	return "APIError"
}

type APIResponse struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
}

func (r APIResponse) Error() string {
	return "APIResponse"
}

func MakeHandler(fn func(ctx *fiber.Ctx) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Call the handler function
		err := fn(ctx)
		if err != nil {
			// Check if the error is an APIError and respond accordingly
			var apiErr APIError
			if errors.As(err, &apiErr) {
				return ctx.Status(apiErr.Code).JSON(apiErr)
			}

			// Check if the error is an APIResponse and respond accordingly
			var apiResp APIResponse
			if errors.As(err, &apiResp) {
				return ctx.Status(apiResp.Code).JSON(apiResp)
			}

			// For other errors, return a generic error response
			log.Error("ISE")
			return ctx.Status(fiber.StatusInternalServerError).JSON(APIError{
				Code: fiber.StatusInternalServerError,
				Err:  "Internal Server Error",
			})
		}
		return nil
	}
}
