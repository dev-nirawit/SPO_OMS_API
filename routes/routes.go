package routes

import (
	"SPO_OMS_API/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.Use(middleware.JWTMiddleware())
	// authRoutes.AuthRoutes(e.Group("/auth"))
}
