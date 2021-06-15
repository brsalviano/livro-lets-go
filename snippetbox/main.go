package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	//Adicionando a checagem se a rota não for exatamente /
	//vamos mandar uma resposta NotFound 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Snippetbox"))
}

// Novos handlers

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mostra um snippet específico..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cria um novo snippet..."))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	//Adicionando handlers no servemux
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Iniciando servidor na porta 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
