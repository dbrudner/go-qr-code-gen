// go-qr-code-gen/internal/ticket/scan.go
package ticket

import (
	"fmt"
	"time"

	"github.com/dbrudner/go-qr-code-gen/internal/db"
)

type Scan struct {
	ID          int
	TicketID    string
	CreatedAt   time.Time
	Fingerprint string
}

func NewScan(ticketID, fingerprint string) (*Scan, error) {
	query := `INSERT INTO scans (ticket_id, fingerprint) VALUES (?, ?)`
	result, err := db.DB.Exec(query, ticketID, fingerprint)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new scan: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return &Scan{
		ID:          int(id),
		TicketID:    ticketID,
		CreatedAt:   time.Now(),
		Fingerprint: fingerprint,
	}, nil
}

func GetScansForTicket(ticketID string) ([]Scan, error) {
	query := `SELECT id, ticket_id, created_at, fingerprint FROM scans WHERE ticket_id = ? ORDER BY created_at DESC`
	rows, err := db.DB.Query(query, ticketID)
	if err != nil {
		return nil, fmt.Errorf("failed to get scans for ticket: %w", err)
	}
	defer rows.Close()

	var scans []Scan
	for rows.Next() {
		var s Scan
		err := rows.Scan(&s.ID, &s.TicketID, &s.CreatedAt, &s.Fingerprint)
		if err != nil {
			return nil, fmt.Errorf("failed to scan scan row: %w", err)
		}
		scans = append(scans, s)
	}

	return scans, nil
}
