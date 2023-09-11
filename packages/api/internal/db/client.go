package db

import (
	"fmt"
	"net/url"
	"os"

	postgrest "github.com/nedpals/postgrest-go/pkg"
)

type DB struct {
	Client *postgrest.Client
}

var (
	supabaseURL = os.Getenv("SUPABASE_URL")
	supabaseKey = os.Getenv("SUPABASE_KEY")
)

func NewClient() (*DB, error) {
	// The /rest/v1/ at the end of url is required
	parsedURL, err := url.Parse(supabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Supabase URL '%s': %w", supabaseURL, err)
	}

	client := postgrest.NewClient(
		*parsedURL,
		postgrest.WithTokenAuth(supabaseKey),
		func(c *postgrest.Client) {
			c.AddHeader("apikey", supabaseKey)
		},
	)

	return &DB{
		Client: client,
	}, nil
}

func (db *DB) Close() {
	db.Client.CloseIdleConnections()
}