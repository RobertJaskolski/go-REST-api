package api

import (
	"github.com/RobertJaskolski/go-REST-api/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"strings"
)

// VALIDATOR

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (s *Server) SetupValidator() {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})
	s.router.Validator = &CustomValidator{validator: validator.New()}
}

// JWT AUTH MIDDLEWARE

func JWTAuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.Envelope{
				"message": "Token is required",
			})
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		userClaims, err := utils.ParseAccessToken(token)
		if err != nil {
			errMsg := "Token is invalid"
			if strings.Contains(err.Error(), "token is expired") {
				errMsg = "Token is expired"
			}
			return echo.NewHTTPError(http.StatusUnauthorized, utils.Envelope{
				"message": errMsg,
			})
		}

		if userClaims.StandardClaims.Valid() != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.Envelope{
				"message": "Token is invalid",
			})
		}

		return next(c)
	}
}
