package handler

import (
	"fmt"

	site "github.com/dbrudner/go-qr-code-gen/internal/site"
	ticketView "github.com/dbrudner/go-qr-code-gen/views/ticket"
	"github.com/labstack/echo/v4"
)

type TicketHandler struct{}

func (h TicketHandler) HandleNewTicket(c echo.Context) error {
	siteID := c.Param("id")
	fmt.Println("hey")
	site, err := site.GetSite(siteID)
	if err != nil {
		fmt.Println("Error")
	}
	return render(c, ticketView.New(*site))
}

func (h TicketHandler) HandleCreateTicket(c echo.Context) error {
	newTicketNotes := c.FormValue("notes")
	fmt.Println(newTicketNotes)
	siteID := c.Param("id")

	site, err := site.GetSite(siteID)
	if err != nil {
		fmt.Println("Error")
	}
	return render(c, ticketView.New(*site))
	// newSite, err := site.NewSite(newSiteDescription, newSiteURL)
	//
	// if err != nil {
	// 	fmt.Println("Error")
	// 	fmt.Println(err)
	// 	return render(c, siteView.New())
	// }
	//
	// fmt.Println(newSiteURL)
	// return render(c, siteView.Created(newSite.URL, newSite.Description))
}
