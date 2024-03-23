package utils

import (
	"fmt"
	"github.com/RobertJaskolski/go-REST-api/config"
	"github.com/RobertJaskolski/go-REST-api/internal/models"
	"github.com/golang-jwt/jwt"
	"os"
)

type JWTUserClaims struct {
	jwt.StandardClaims
	ID    int                 `json:"id"`
	Email string              `json:"email"`
	Role  models.UserRoleEnum `json:"role"`
}

func CreateJWTToken(claims JWTUserClaims, cfg *config.Config) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

func CreateJWTRefreshToken(claims jwt.StandardClaims, cfg *config.Config) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.RefreshSecret))
}

func ParseAccessToken(accessToken string) (*JWTUserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &JWTUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return parsedAccessToken.Claims.(*JWTUserClaims), err
}

func ParseRefreshToken(refreshToken string) (*jwt.StandardClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})

	fmt.Println(err)

	return parsedRefreshToken.Claims.(*jwt.StandardClaims), err
}
