package handlers

import (
	"context"
	"github.com/RobertJaskolski/go-REST-api/config"
	"github.com/RobertJaskolski/go-REST-api/internal/models"
	"github.com/RobertJaskolski/go-REST-api/internal/repositories"
	"github.com/RobertJaskolski/go-REST-api/internal/utils"
	"github.com/golang-jwt/jwt"
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

	// Check if user exists
	user, err := handler.UserRepository.GetLoggedByEmail(context.Background(), dto.Email)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, utils.Envelope{"message": err.Error()})
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, utils.Envelope{"message": "user or password are incorrect"})
	}

	if !user.IsActive {
		return ctx.JSON(http.StatusUnauthorized, utils.Envelope{"message": "user is not active"})
	}

	// Generate JWT token (access token)
	token, err := utils.CreateJWTToken(utils.JWTUserClaims{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		},
	}, handler.cfg)

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, utils.Envelope{"message": "An error occurred while generating token"})
	}

	// Generate JWT token (refresh token)
	refreshToken, err := utils.CreateJWTRefreshToken(jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
	}, handler.cfg)

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, utils.Envelope{"message": "An error occurred while generating token"})
	}

	return ctx.JSON(http.StatusOK, utils.Envelope{
		"access_token":       token,
		"refresh_token":      refreshToken,
		"expires_in":         time.Now().Add(time.Hour * 6).Unix(),
		"refresh_expires_in": time.Now().Add(time.Hour * 48).Unix(),
	})
}
