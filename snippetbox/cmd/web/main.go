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

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Inicializa uma nova struct http.Server. Nós configuramos o Addr e o campo Handler
	//então o servidor vai usar o mesmo endereço da rede e rota que antes, além disso,
	//nós configuramos o campo ErrorLog, então o servidor vai passar a usar o logger personalizado errorLog
	//toda vez que tiver algum problema.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Servidor escutando na porta %s", *addr)

	//A partir de agora vamos chamar o método ListenAndServe da instância de http.Server personalizada.
	err := srv.ListenAndServe()
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
