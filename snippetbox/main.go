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
		//Vamos enviar o header Allow: POST no header map da resposta.
		//O primeiro parâmetro é o nome do cabeçalho e o segundo é seu valor.
		w.Header().Set("Allow", http.MethodPost)
		//Importante: Usar w.Header().Set() depois de w.WriteHeader ou w.Write
		//não modifica o header que o usuário vai receber. Portanto, qualquer
		//modificação no header precisa ser informada antes.
		w.WriteHeader(405)
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
