package main

import (
	"html/template"
	"io"

	"github.com/dbrudner/go-qr-code-gen/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	t := &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	return t
}

type PageData struct {
	Site string
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/public", "./public")
	db.Init("db.sqlite")
	db.CreateTables()
	// db.SeedData()
	pageData := PageData{Site: "Mysite.com"}
	e.Renderer = newTemplates()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", pageData)
	})

	e.Logger.Fatal(e.Start(":3005"))
}
