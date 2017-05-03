package controllers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mucyomiller/BambooAPI/models"
)

/*
*
*      Configurations for rendering html templates
*
 */

//Template struct
type (
	Template struct {
		templates *template.Template
	}
)

//T var of type Template
var T = &Template{
	templates: template.Must(template.ParseGlob("templates/*/*.gohtml")),
}

//Render func for our Template T
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

/*
*
*    API functions (return json response)
*
 */

//GetAll func get All Menus
func GetAll(c echo.Context) error {
	ms, _ := models.All()

	return c.JSON(http.StatusOK, ms)
}

//GetByID func get Menu By ID
func GetByID(c echo.Context) error {
	id := c.Param("id")

	m, _ := models.GetByID(id)

	return c.JSON(http.StatusOK, m)
}

//NewMenu func create New Menu
func NewMenu(c echo.Context) error {
	m := new(models.Menu)

	c.Bind(m)

	nm, _ := models.NewMenu(*m)

	return c.JSON(http.StatusOK, nm)
}

//RemoveMenu func remove Menu
func RemoveMenu(c echo.Context) error {
	id := c.Param("id")

	models.DeleteMenu(id)
	return c.String(http.StatusOK, "")
}

/*
*
* function for dealing with views(return rendered html)
*
 */

//Index func
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}

//Home func
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Miller!")
}

//Create menu func for testing mongo insert
func Create(c echo.Context) error {
	m := models.Menu{Name: "dojos", Description: "Some water menus", Price: 10000.0, ImageURL: "http://localhost/images/test.jpeg"}
	m, err := models.NewMenu(m)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, "Ok!")
}
