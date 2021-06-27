package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	//Estamos extraindo o valor do parâmetro id da query string
	//e tentando converte-lo em um inteiro através de strconv.Atoi().
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	//Se não puder ser convertido em um inteiro ou se o valor for menor que 1,
	//vamos retornar uma página de erro 404 na resposta.
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//A função Fprintf pode ser usada no lugar de w.Write porque ela pede
	//que o primeiro parâmetro seja um io.writer e o parâmetro w corresponde a essa interface.
	fmt.Fprintf(w, "Mostra um snippet específico com o ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
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
