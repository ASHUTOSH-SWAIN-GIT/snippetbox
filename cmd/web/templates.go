package main

import "snippetbox.ashutosh.net/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
