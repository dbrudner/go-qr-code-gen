package ticket

import (
	"fmt"

	"github.com/dbrudner/go-qr-code-gen/internal/db"
	"github.com/google/uuid"
)

func NewTicket(siteId string, userId string, content string) (*Ticket, error) {
	id := uuid.New()

	query := `INSERT INTO tickets (id, site_id, user_id, content) VALUES (?, ?, ?, ?)`
	_, err := db.DB.Exec(query, id, siteId, userId, content)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new ticket: %w", err)
	}

	return &Ticket{}, nil
}

func GetTicket(id string) (*Ticket, error) {
	query := `SELECT id, site_id, user_id, content FROM tickets WHERE id = $1;`
	row := db.DB.QueryRow(query, id)

	var t Ticket
	err := row.Scan(&t.ID, &t.Site_ID, &t.User_ID, &t.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to get ticket: %w", err)
	}

	return &t, nil
}
