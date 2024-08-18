package main

import (
	"html/template"

	handler "github.com/dbrudner/go-qr-code-gen/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/public", "./public")
	// db.Init("db.sqlite")
	// db.CreateTables()
	// db.SeedData()

	homeHandler := handler.HomeHandler{}
	e.GET("/", homeHandler.HandleHomeShow)
	e.Logger.Fatal(e.Start(":3005"))
}
