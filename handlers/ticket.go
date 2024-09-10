package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"

	site "github.com/dbrudner/go-qr-code-gen/internal/site"
	"github.com/dbrudner/go-qr-code-gen/internal/ticket"
	ticketView "github.com/dbrudner/go-qr-code-gen/views/ticket"
)

type TicketHandler struct{}

// For handling form/POST route
func (h TicketHandler) HandleCreateTicket(c echo.Context) error {
	newTicketContent := c.FormValue("content")
	fmt.Println(newTicketContent, "content")
	siteID := c.Param("id")
	fmt.Println(siteID, "siteID")

	// Check if site exists -- this should probalby be middleware driven
	_, err := site.GetSite(siteID)
	if err != nil {
		fmt.Println("What the heck?")
		return err
	}

	newTicket, err := ticket.NewTicket(
		siteID, 1, newTicketContent)
	if err != nil {
		fmt.Println("Error creating ticket")
		return err
	}

	fmt.Printf("Created new ticket: %s", newTicket.ID)
	return c.Redirect(301, fmt.Sprintf("/site/%s/ticket/%s", siteID, newTicket.ID))
}

func (h TicketHandler) HandleTicketDetail(c echo.Context) error {
	siteID := c.Param("id")
	site, err := site.GetSite(siteID)
	if err != nil {
		fmt.Println("Error finding site")
	}

	ticketID := c.Param("ticketId")
	fmt.Println(ticketID, "ticketID")
	if ticketID == "new" {
		return render(c, ticketView.New(*site))
	}

	ticket, err := ticket.GetTicket(ticketID)
	if err != nil {
		fmt.Println("Error finding ticket")
	}

	fmt.Println("rendering detail view")
	return render(c, ticketView.Detail(*ticket, site.URL))
}
