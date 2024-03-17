package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/RobertJaskolski/go-REST-api/internal/models"
	"github.com/RobertJaskolski/go-REST-api/internal/repositories"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	UserRepository repositories.UserRepository
}

func NewUserHandler(userRepository repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

func (handler *UserHandler) CreateUser(ctx echo.Context) error {
	dto := new(models.CreateUserDTO)
	if err := ctx.Bind(dto); err != nil {
		fmt.Println("BIND ERROR")
		return err
	}

	if err := ctx.Validate(dto); err != nil {
		return err
	}

	generatePassword := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(generatePassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	dto.Password = string(hashedPassword)

	entityID, err := handler.UserRepository.Create(context.Background(), dto)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return ctx.JSON(http.StatusBadRequest, pgErr.Detail)
		}
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(201, entityID)
}

func (handler *UserHandler) GetUsers(ctx echo.Context) error {
	return nil
}

func (handler *UserHandler) GetUser(ctx echo.Context) error {
	return nil
}

func (handler *UserHandler) UpdateUser(ctx echo.Context) error {
	return nil
}

func (handler *UserHandler) DeleteUser(ctx echo.Context) error {
	return nil
}
