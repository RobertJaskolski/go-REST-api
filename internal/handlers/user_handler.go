package handlers

import (
	"context"
	"github.com/RobertJaskolski/go-REST-api/internal/models"
	"github.com/RobertJaskolski/go-REST-api/internal/repositories"
	"github.com/RobertJaskolski/go-REST-api/internal/utils"
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
	err := utils.Validate(ctx, dto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	generatePassword := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(generatePassword), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}
	dto.Password = string(hashedPassword)

	user, err := handler.UserRepository.Create(context.Background(), dto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (handler *UserHandler) GetUsers(ctx echo.Context) error {
	return nil
}

func (handler *UserHandler) GetUser(ctx echo.Context) error {
	id, err := utils.GetID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	user, err := handler.UserRepository.GetOne(context.Background(), id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)
}

func (handler *UserHandler) UpdateUser(ctx echo.Context) error {
	id, err := utils.GetID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	dto := new(models.UpdateUserDTO)

	err = utils.Validate(ctx, dto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	user, err := handler.UserRepository.Update(context.Background(), id, dto)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	return ctx.JSON(http.StatusAccepted, user)
}

func (handler *UserHandler) DeleteUser(ctx echo.Context) error {
	id, err := utils.GetID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	err = handler.UserRepository.Delete(context.Background(), id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	return ctx.JSON(http.StatusAccepted, utils.Envelope{"success": true})
}
