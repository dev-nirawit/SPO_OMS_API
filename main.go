package main

import (
	"SPO_OMS_API/config"
	"SPO_OMS_API/middleware"
	"SPO_OMS_API/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDB()

	// สร้าง Echo instance
	e := echo.New()

	// กำหนด middleware
	e.Use(middleware.JWTMiddleware())

	// กำหนดเส้นทาง
	routes.InitRoutes(e)
	// เริ่ม server
	e.Logger.Fatal(e.Start(":8099"))
}
