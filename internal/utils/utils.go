package utils

import (
	"github.com/labstack/echo/v4"
)

type Envelope map[string]interface{}

func Validate(ctx echo.Context, dto interface{}) error {
	if err := ctx.Bind(dto); err != nil {
		return err
	}

	if err := ctx.Validate(dto); err != nil {
		return err
	}

	return nil
}
