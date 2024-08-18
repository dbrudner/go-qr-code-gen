package site

import (
	"fmt"

	"github.com/dbrudner/go-qr-code-gen/internal/db"
	"github.com/google/uuid"
)

func NewSite(description, url string) (*Site, error) {
	id := uuid.New()

	query := `INSERT INTO sites (id, description, url) VALUES (?, ?, ?)`
	_, err := db.DB.Exec(query, id, description, url)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new site: %w", err)
	}

	return &Site{
		ID:          id.String(),
		Description: description,
		URL:         url,
	}, nil
}
