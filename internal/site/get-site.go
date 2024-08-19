package site

import (
	"fmt"

	"github.com/dbrudner/go-qr-code-gen/internal/db"
)

func GetSite(id string) (*Site, error) {
	query := `SELECT id, description, url FROM sites WHERE id = $1;`
	row := db.DB.QueryRow(query, id)

	var s Site
	err := row.Scan(&s.ID, &s.Description, &s.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to get site: %w", err)
	}

	return &s, nil
}
