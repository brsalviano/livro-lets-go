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
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)
	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return s, nil

}

//Latest retorna os 10 snippets mais recentes
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {

	//Escreve a declaração SQL que desejamos executar
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	// Usa o método Query no pool de conexões para executar a declaração SQL.
	// Retorna um sql.Rows que é um resultset contendo os resultados da query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// Aqui vamos garantir que sql.Rows é fechado apropriadamente, depois de todas as instruções.
	// Esta declaração defer deve vir depois você verificar os erros do método Query.
	// Se colocarmos antes e o método Query retornar um erro, vamos receber um panic por estarmos tentando
	// encerrar uma conexão de um resultset nil.
	defer rows.Close()

	//Inicializa um slice vazio para guardar objetos models.Snippets.
	snippets := []*models.Snippet{}

	// Usa rows.Next para iterar pelas linhas do resultset.
	// Prepara a primeira (e então subsequente) linhas para ser atuado pelo método
	// rows.Scan(). Se a iteração por todas as linhas terminar, então o resultset é
	// automaticamente encerrado e liberado da conexão do banco de dados.
	for rows.Next() {
		//Cria um ponteiro para uma struct Snippet zerada.
		s := &models.Snippet{}

		// Usa rows.Scan() para copiar os valores de cada campo da linha para
		// o novo objeto Snippet que criamos. Novamente, os argumentos para row.Scan()
		// devem ser pointeiros para os locais que desejamos copiar os dados, e o número
		// de argumentos deve ser exatamente igual ao número  de colunas retornadas da declaração
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		// Acrescenta o Snippet recém criado e configurado com os dados da linha atual no slice de Snippets
		snippets = append(snippets, s)
	}

	// Quando o loop rows.Next é concluído nós chamamos rows.Err() para obter qualquer erro encontrado durante
	// a iteração. É importante chamar isso - Não assuma que a iteração bem-sucedida foi completada por todo
	// o resultset.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	//Se tudo ocorreu bem, vamos retornar o slice de Snippets
	return snippets, nil

}
