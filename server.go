package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mucyomiller/BambooAPI/controllers"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", controllers.Index)
	//e.GET("/favicon.ico", echo.WrapHandler(http.NotFoundHandler()))
	e.Logger.Fatal(e.Start(":8080"))
}
