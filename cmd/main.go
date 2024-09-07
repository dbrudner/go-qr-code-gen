package main

import (
	"html/template"

	handler "github.com/dbrudner/go-qr-code-gen/handlers"
	db "github.com/dbrudner/go-qr-code-gen/internal/db"
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
	db.Init("db.sqlite")
	db.CreateTables()
	db.SeedData()

	homeHandler := handler.HomeHandler{}
	sitesHandler := handler.SiteHandler{}
	ticketHandler := handler.TicketHandler{}

	e.GET("/", homeHandler.HandleHomeShow)
	e.GET("/sites", sitesHandler.HandleSiteCollection)
	e.GET("/site/:id", sitesHandler.HandleSiteDetail)
	e.GET("/site/new", sitesHandler.HandleNewSite)
	e.GET("/site/:id/ticket/new", ticketHandler.HandleNewTicket)
	e.GET("site/:id/ticket/:ticketId", ticketHandler.HandleTicketDetail)
	e.POST("/site/new", sitesHandler.HandleCreateSite)
	e.POST("/site/:id/ticket/new", ticketHandler.HandleCreateTicket)

	e.Logger.Fatal(e.Start(":3005"))
}
