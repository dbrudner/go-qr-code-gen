package ticket

import (
	"fmt"
	"time"

	"github.com/dbrudner/go-qr-code-gen/internal/db"
	"github.com/google/uuid"
)

type Ticket struct {
	ID        string
	SiteID    string
	UserID    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Scans     []Scan
}

func NewTicket(siteId string, userId int, content string) (*Ticket, error) {
	id := uuid.New()

	query := `INSERT INTO tickets (id, site_id, user_id, content) VALUES (?, ?, ?, ?)`
	_, err := db.DB.Exec(query, id, siteId, userId, content)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new ticket: %w", err)
	}

	return &Ticket{
		ID:      id.String(),
		SiteID:  siteId,
		UserID:  userId,
		Content: content,
	}, nil
}

func GetTicket(id string) (*Ticket, error) {
	query := `SELECT id, site_id, user_id, content, created_at, updated_at FROM tickets WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var t Ticket
	err := row.Scan(&t.ID, &t.SiteID, &t.UserID, &t.Content, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get ticket: %w", err)
	}

	scans, err := GetScansForTicket(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get scans for ticket: %w", err)
	}

	t.Scans = scans

	return &t, nil
}
