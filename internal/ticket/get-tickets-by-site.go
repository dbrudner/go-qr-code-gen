package ticket

import (
	"fmt"

	"github.com/dbrudner/go-qr-code-gen/internal/db"
	"github.com/dbrudner/go-qr-code-gen/internal/site"
)

type SiteWithTickets struct {
	Site    site.Site
	Tickets []Ticket
}

func GetSiteWithTickets(siteID string) (*SiteWithTickets, error) {
	fmt.Println(siteID, " sitei1")
	query := `
	SELECT
		sites.id,
		sites.url,
		tickets.id AS ticket_id,
		tickets.user_id,
		tickets.created_at,
		tickets.updated_at AS ticket_updated_at,
		tickets.content
	FROM
		sites
	LEFT JOIN
		tickets ON sites.id = tickets.site_id
	WHERE
		sites.id = ?
	`
	rows, err := db.DB.Query(query, siteID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	if rows == nil {
		return nil, fmt.Errorf("rows is nil")
	}
	defer rows.Close()

	var site SiteWithTickets
	var tickets []Ticket

	for rows.Next() {
		var ticket Ticket
		fmt.Println("site id")
		fmt.Println(&site.Site.ID)
		err := rows.Scan(
			&site.Site.ID,
			&site.Site.URL,
			&ticket.ID,
			&ticket.UserID,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
			&ticket.Content,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		tickets = append(tickets, ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	site.Tickets = tickets
	fmt.Println("ticket length", len(site.Tickets))
	return &site, nil
}
