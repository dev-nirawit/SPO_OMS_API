package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginHandler จัดการการเข้าสู่ระบบ
func LoginHandler(c echo.Context) error {
	// Logic สำหรับการเข้าสู่ระบบ
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
	})
}

// ProfileHandler จัดการการดึงข้อมูลโปรไฟล์
func ProfileHandler(c echo.Context) error {
	// Logic สำหรับการดึงข้อมูลโปรไฟล์
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": "User details",
	})
}
