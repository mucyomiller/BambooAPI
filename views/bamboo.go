package views

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type (
	Template struct {
		templates *template.Template
	}
)

var T = &Template{
	templates: template.Must(template.ParseGlob("templates/*/*.gohtml")),
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Miller!")
}
