package utils

import (
	"github.com/RobertJaskolski/go-REST-api/config"
	"github.com/RobertJaskolski/go-REST-api/internal/models"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateJWTToken(user *models.LoggedUser, cfg *config.Config) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(cfg.JWT.Secret))
}

func CreateJWTRefreshToken(user *models.LoggedUser, cfg *config.Config) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte(cfg.JWT.RefreshSecret))
}
