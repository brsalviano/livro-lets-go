package mysql

import (
	"database/sql"

	"github.com/brsalviano/livro-lets-go/snippetbox/pkg/models"
)

//SnippetModel embrulha o pool de conexões sql.DB
type SnippetModel struct {
	DB *sql.DB
}

//Insert insere um novo snippet no banco de dados
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// Escreve a declaração SQL que desejamos executar. Eu dividi em duas linhas
	// por questões de legibilidade (motivo pelo qual usei backtick em vez de aspas duplas)
	// Os placeholders (?) são usados para proteger o código contra sql injection.
	statement := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Usa o método Exec no pool de conexões incorporado para executar a declaração.
	// O primeiro parâmetro é a declaração SQL, seguido pelo título, conteúdo e valor de expiração
	// para o placeholder de parâmetros (?). Este método retorna um objeto sql.Result, que contém
	// informações básicas sobre o que aconteceu quando a declaração foi executada.
	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Usa o método LastInsertId no resultado do objeto para obter o ID,
	// do registro da tabela snippet recém criado.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// O ID retornado tem o tipo int64, então precisamos converter para o tipo int
	// antes de retornar.
	return int(id), nil

}

//Get retorna um snippet específico baseado no id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

//Latest retorna os 10 snippets mais recentes
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
