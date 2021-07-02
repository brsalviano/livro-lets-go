package mysql

import (
	"database/sql"

	//Importa o pacote de modelos que acabamos de criar.
	//Necessário passar o caminho do módulo (nome do módulo conforme está em go.mod)
	"github.com/brsalviano/livro-lets-go/snippetbox/pkg/models"
)

//SnippetModel embrulha o pool de conexões sql.DB
type SnippetModel struct {
	DB *sql.DB
}

//Insert insere um novo snippet no banco de dados
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

//Get retorna um snippet específico baseado no id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

//Latest retorna os 10 snippets mais recentes
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
