package handlers

import (
	"context"
	"github.com/RobertJaskolski/go-REST-api/config"
	"github.com/RobertJaskolski/go-REST-api/internal/models"
	"github.com/RobertJaskolski/go-REST-api/internal/repositories"
	"github.com/RobertJaskolski/go-REST-api/internal/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthHandler struct {
	cfg            *config.Config
	UserRepository repositories.UserRepository
}

func NewAuthHandler(cfg *config.Config, userRepository repositories.UserRepository) *AuthHandler {
	return &AuthHandler{
		UserRepository: userRepository,
		cfg:            cfg,
	}
}

func (handler *AuthHandler) Login(ctx echo.Context) error {
	dto := new(models.LoginDTO)
	err := utils.Validate(ctx, dto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	user, err := handler.UserRepository.GetLoggedByEmail(context.Background(), dto.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": "user or password are incorrect"})
	}

	token, err := utils.CreateJWTToken(user, handler.cfg)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": "user or password are incorrect"})
	}

	refreshToken, err := utils.CreateJWTToken(user, handler.cfg)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Envelope{"message": "user or password are incorrect"})
	}

	return ctx.JSON(http.StatusOK, utils.Envelope{
		"access_token":       token,
		"refresh_token":      refreshToken,
		"expires_in":         time.Now().Add(time.Hour * 24).UnixMilli(),
		"refresh_expires_in": time.Now().Add(time.Hour * 24 * 7).UnixMilli(),
	})
}
