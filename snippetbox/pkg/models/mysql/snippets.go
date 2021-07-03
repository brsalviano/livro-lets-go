package mysql

import (
	"database/sql"
	"errors" //Novo import

	"github.com/brsalviano/livro-lets-go/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

//Insert insere um novo snippet no banco de dados
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	statement := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

//Get retorna um snippet específico baseado no id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	// Escreve a declaração SQL que desejamos executar. Novamente, o código
	// foi feito em duas linhas por questões de legibilidade.
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Usa o método QueryRow no pool de conexões para executar a declaração SQL,
	// passando a variável não confiável id como valor do placeholder.
	// A declaração retorna um objeto sql.Row que guarda o resultado do banco de dados.
	row := m.DB.QueryRow(stmt, id)

	//Inicializa um ponteiro para uma nova instância de Snippet
	s := &models.Snippet{}

	// Usa row.Scan para copiar os valores de cada campo de sql.Row para o campo correspondente
	// da struct Snippet. Note que os argumentos para row.Scan são ponteiros para colocar o que
	// os dados que você quer copiar para ele, e o número de argumentos deve ser exatamente o número
	// de colunas retornadas pela declaração.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// Se a query não retornar linhas, então row.Scan retorna um erro sql.ErroNoRows.
		// Estamos usando erros.Is para verificar o erro específico e retornar nosso próprio
		// erro models.ErrNoRecord em vez disso.
		if errors.Is(err, sql.ErrNoRows) {
			// A ideia de ter nosso pŕoprio ErrNoRecord é encapsular completamente o nosso modelo.
			// Dessa forma, nossa aplicação não se preocupa com erros específicos de um determinado
			// driver etc.
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return s, nil

}

//Latest retorna os 10 snippets mais recentes
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
