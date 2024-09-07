package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"

	site "github.com/dbrudner/go-qr-code-gen/internal/site"
	"github.com/dbrudner/go-qr-code-gen/internal/ticket"
	"github.com/dbrudner/go-qr-code-gen/views/home"
	ticketView "github.com/dbrudner/go-qr-code-gen/views/ticket"
)

type TicketHandler struct{}

// for handling page GET route
func (h TicketHandler) HandleNewTicket(c echo.Context) error {
	siteID := c.Param("id")
	fmt.Println("hey")
	site, err := site.GetSite(siteID)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(site.GetURL())
	return render(c, ticketView.New(*site))
}

// For handling form/POST route
func (h TicketHandler) HandleCreateTicket(c echo.Context) error {
	newTicketNotes := c.FormValue("notes")
	fmt.Println(newTicketNotes)
	siteID := c.Param("siteId")

	// Check if site exists -- this should probalby be middleware driven
	_, err := site.GetSite(siteID)
	if err != nil {
		fmt.Println("What the heck?")
	}

	newTicket, err := ticket.NewTicket(
		siteID, "my-user-id", newTicketNotes)
	if err != nil {
		fmt.Println("Error creating ticket")
		return err
	}

	fmt.Printf("Created new ticket: %s", newTicket.ID)
	return render(c, home.Show())
	// return render(c, ticketView.New(*site))
	// return render(c, ticketView.New(*site))
}

func (h TicketHandler) HandleTicketDetail(c echo.Context) error {
	siteID := c.Param("id")
	ticketID := c.Param("ticketId")
	fmt.Println(siteID)
	fmt.Println(ticketID)
	site, err := site.GetSite(siteID)
	if err != nil {
		fmt.Println("Error finding site")
	}

	return render(c, ticketView.Detail(site.URL))
}
