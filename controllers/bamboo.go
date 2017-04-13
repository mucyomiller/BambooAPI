package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mucyomiller/BambooAPI/models"
)

//GetAll func get All Menus
func GetAll(c echo.Context) error {
	ms, _ := models.All()

	return c.JSON(http.StatusOK, ms)
}

//GetByID func get Menu By ID
func GetByID(c echo.Context) error {
	id := c.Param("id")

	m, _ := models.GetById(id)

	return c.JSON(http.StatusOK, m)
}

//NewMenu func create New Menu
func NewMenu(c echo.Context) error {
	m := new(models.Menu)

	c.Bind(m)

	nm, _ := models.NewMenu(*m)

	return c.JSON(http.StatusOK, nm)
}

//RemoveMenu func remove Mene
func RemoveMenu(c echo.Context) error {
	id := c.Param("id")

	models.DeleteMenu(id)
	return c.String(http.StatusOK, "")
}
