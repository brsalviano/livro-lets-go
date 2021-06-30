package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Estamos passando a nossa implementação do FileSystem como parâmetro de FileServer
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Iniciando servidor na porta 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Criando uma struct para encapsular nossa implementação personalizada do http file system
type neuteredFileSystem struct {
	fs http.FileSystem
}

// Para corresponder a interface FileSystem criamos uma função associada chamada
// Open que recebe como parâmetro o caminho até o arquivo e retorna o arquivo ou um erro
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {

	//Tenta abrir o diretório/arquivo com a implementação do sistema de arquivos
	//que nós encapsulamos (http.FileSystem)
	f, err := nfs.fs.Open(path)
	if err != nil {
		//Se der erro, eu já retorno sem o arquivo e com o erro
		return nil, err
	}

	// Vou tentar pegar as informações do arquivo
	s, err := f.Stat()
	//Se for uma pasta
	if s.IsDir() {
		//Faço a união do caminho passado com index.html
		index := filepath.Join(path, "index.html")
		//Faço um if com bloco de declaração tentando abrir o caminho com o arquivo index.html
		if _, err := nfs.fs.Open(index); err != nil {
			//Se der erro é porque o arquivo não existe
			//então eu fecho o arquivo que está aberto no sistema operacional até agora...
			closeErr := f.Close()
			if closeErr != nil {
				//Se eu não conseguir fechar eu retorno sem o arquivo e com o erro
				return nil, closeErr
			}

			//Mesmo que eu feche, como o arquivo não existe eu retorno sem o arquivo e com o erro
			return nil, err
		}
	}

	//Se chegar até aqui é porque o arquivo pode não ser uma pasta, então eu retorno o próprio arquivo.
	//Se for uma pasta e chegou aqui, é porque tem o arquivo index.html, então não tem problema,
	//já que este arquivo será executado evitando que a listagem seja feita
	return f, nil
}
