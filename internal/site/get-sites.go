package site

import (
	"fmt"

	"github.com/dbrudner/go-qr-code-gen/internal/db"
)

func GetSites() ([]Site, error) {
	query := `SELECT id, description, url FROM sites;`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get sites: %w", err)
	}

	defer rows.Close()

	var sites []Site
	for rows.Next() {
		var s Site
		err = rows.Scan(&s.ID, &s.Description, &s.URL)
		if err != nil {
			return nil, err
		}
		sites = append(sites, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sites, nil
}
