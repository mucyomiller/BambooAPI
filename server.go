package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mucyomiller/BambooAPI/controllers"
)

func main() {

	e := echo.New()
	e.Debug = true
	e.Renderer = controllers.T
	// Middleware
	// logging request info
	e.Use(middleware.Logger())
	//recovering from failure
	e.Use(middleware.Recover())
	// gzipping content to reduce payload size
	e.Use(middleware.Gzip())

	e.Static("/static", "assets")
	e.GET("/", controllers.Index)
	e.GET("/home", controllers.Home)
	e.GET("/create", controllers.Create)
	//e.GET("/favicon.ico", echo.WrapHandler(http.NotFoundHandler()))
	e.Logger.Fatal(e.Start(":8080"))
}
