package main

import (
	"database/sql" //importando o pacote nativo para trabalhar com sql
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" //Importando e carregando o driver
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	//Definindo uma nova flag de linha de comando para a string DSN do mysql
	dsn := flag.String("dsn", "root:teste@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Para manter a função main limpa. Eu coloquei o código para criar o pool de conexões
	//em uma função separada openDB(). Nós passamos a DSN da flag de linha de comando.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Nós também usamos defer com db.Close para a conexão ser
	// encerrada antes da função main ser encerrada.
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Servidor escutando na porta %s", *addr)

	// Como a variável err já está declarada acima, modificamos o código
	// para usar o operador de atribuição = em vez do operador de declaração e atribuição :=
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// A função openDB, encapsula sql.Open() e retorna um pool de conexões sql.DB
// para a DSN fornecida.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//Estamos usando db.Ping para ver se a conexão pode ser estabelecida.
	if err = db.Ping(); err != nil {
		return nil, err
	}
	//Se não der nenhum problema no ping, retornamos *sql.DB
	return db, nil
}
