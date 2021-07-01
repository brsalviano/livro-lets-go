package main

import (
	"flag" //novo import
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	//Define uma nova flag de linha de comando com o nome 'addr',
	//com o valor padrão ':4000' e um texto curto explicando o que a flag controla.
	//O valor da flag será guardado na variável addr em tempo de execução.
	addr := flag.String("addr", ":4000", "HTTP network address")

	//Importante. Nós usamos flag.Parse para fazer a análise da flag da linha de comando.
	//flag.Parse faz a leitura do valor da flag da linha de comando e atribui para a variável addr.
	//É necessário chamar o Parse antes de usar a variável senão sempre vai conter o valor
	//padrão :4000. Se algum erro for encontrado durante a análise a aplicação será encerrada.
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// O valor retornado de flag.String() é um ponteiro para o valor da flag, não o valor propriamente dito.
	// Então nós precisamos desreferenciar o ponteiro antes de usá-lo (por isso o *).
	log.Printf("Servidor escutando na porta %s", *addr)

	//Agora em vez de passar uma porta fixa, passamos o valor recebido da flag *addr
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
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
