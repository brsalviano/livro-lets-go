package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mostra um snippet específico..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	//Verificando se o método da requisição é POST
	if r.Method != http.MethodPost {
		//O status será 405 (Method Not Allowed)
		w.WriteHeader(405)
		//A tela vai mostrar a mensagem especificada abaixo
		w.Write([]byte("Método não permitido"))
		return
	}

	w.Write([]byte("Cria um novo snippet..."))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Iniciando servidor na porta 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
