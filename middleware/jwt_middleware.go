package middleware

import (
	"SPO_OMS_API/config"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware สร้าง JWT middleware
func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JwtSecret), // Secret key สำหรับ JWT
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  false,
				"message": "missing or malformed jwt",
				"data":    []interface{}{},
			})
		},
	})
}
