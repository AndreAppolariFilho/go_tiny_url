package main

import (
	"time"

	"github.com/AndreAppolariFilho/go_tiny_url/internal/database"
	"github.com/google/uuid"
)

type Url struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OriginalUrl string `json:"original_url"`
	TinyUrl     string `json:"tiny_url"`
}

func databaseUrlToUrl(dbUrl database.Url) Url{
	return Url{
		ID: dbUrl.ID,
		CreatedAt: dbUrl.CreatedAt,
		UpdatedAt: dbUrl.UpdatedAt,
		OriginalUrl: dbUrl.OriginalUrl,
		TinyUrl: dbUrl.TinyUrl,
	}
}
