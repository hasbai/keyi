package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HttpError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Detail  *ErrorDetail `json:"detail"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func BadRequest(messages ...string) *HttpError {
	message := "Bad Request"
	if len(messages) > 0 {
		message = messages[0]
	}
	return &HttpError{
		Code:    400,
		Message: message,
	}
}

func Unauthorized(messages ...string) *HttpError {
	message := "Invalid JWT Token"
	if len(messages) > 0 {
		message = messages[0]
	}
	return &HttpError{
		Code:    401,
		Message: message,
	}
}

func Forbidden(messages ...string) *HttpError {
	message := "您没有权限进行此操作"
	if len(messages) > 0 {
		message = messages[0]
	}
	return &HttpError{
		Code:    403,
		Message: message,
	}
}

func MyErrorHandler(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	httpError := HttpError{
		Code:    500,
		Message: err.Error(),
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		httpError.Code = 404
	} else {
		switch e := err.(type) {
		case *HttpError:
			httpError = *e
		case *fiber.Error:
			httpError.Code = e.Code
		case *ErrorDetail:
			httpError.Code = 400
			httpError.Detail = e
		case fiber.MultiError:
			httpError.Code = 400
			httpError.Message = ""
			for _, err = range e {
				httpError.Message += err.Error() + "\n"
			}
		}
	}

	//if httpError.Code == 500 && config.Config.Debug {
	//	// log error
	//	debug.PrintStack()
	//}

	return ctx.Status(httpError.Code).JSON(&httpError)
}
