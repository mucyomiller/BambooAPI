package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/mucyomiller/BambooAPI/controllers"
	"github.com/mucyomiller/BambooAPI/views"
)

func main() {

	e := echo.New()
	e.Debug = true
	e.Renderer = views.T
	// Middleware
	// logging request info
	e.Use(middleware.Logger())
	//recovering from failure
	e.Use(middleware.Recover())
	// gzipping content to reduce payload size
	e.Use(middleware.Gzip())

	e.Static("/static", "assets")
	e.GET("/", views.Index)
	e.GET("/home", views.Home)
	//e.GET("/favicon.ico", echo.WrapHandler(http.NotFoundHandler()))
	e.Logger.Fatal(e.Start(":8080"))
}
