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
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		//http.Error é um atalho para WriteHead + Write
		//Nos casos em que queremos especificar algum status diferente de 200.
		http.Error(w, "Método não permitido", 405)
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
