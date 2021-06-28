package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Cria um servidor de arquivos que serve arquivos da pasta ./ui/static
	// Atente-se para o fato de que o caminho passado para http.Dir é relativo a raiz do projeto.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//mux.Handle() está registrando o servidor de arquivos
	//como um handler para as URLs que comecem com /static/.
	//Para a correta correspondência dos caminhos, nós removemos o prefixo
	//"/static" antes da requisição chegar no servidor de arquivos.
	//Senão o caminho seria lido da forma ./ui/static/static/...
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Iniciando servidor na porta 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
