package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	//Usa o log.New() para criar um logger para escrever mensagens informativas.
	//Recebe 3 parâmetros: O destino para escrever os logs (os.Stout), uma string
	//com a mensagem prefixada com a palavra INFO seguido de tab, e flags para indicar
	//quais informações adicionais incluir (local date e time). Perceba que as flags são
	//unidas usando o operador bitwise OR: |
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	//Cria um logger para escrever mensagens de erro da mensma maneira, mas usa Stderr
	//como destino e usa a flag log.Lshortfile para incluir o nome do arquivo relevante e o número da linha.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Escreve as mensagens usando os 2 novos loggers.
	infoLog.Printf("Servidor escutando na porta %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {

	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
