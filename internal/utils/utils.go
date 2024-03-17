package utils

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
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

func GetID(ctx echo.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Invalid ID: %s", ctx.Param("id")))
	}

	return id, nil
}
