package main

import "github.com/brsalviano/livro-lets-go/snippetbox/pkg/models"

// Define um tipo templateData que serve para encapsular
// os dados dinâmicos que queremos passar para os nossos templates HTML.
type templateData struct {
	Snippet *models.Snippet
}
