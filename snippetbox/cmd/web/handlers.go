package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//Como vamos precisar usar mais de um arquivo de template
	//vamos criar um slice especificando cada um deles.
	files := []string{
		"./ui/html/home.page.gohtml", //Aqui a ordem importa!
		"./ui/html/base.layout.gohtml",
	}
	//Na hora de passar para a analise de templates
	//passamos o slice com os templates, lendo-os como variadic .
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Erro interno no servidor", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Erro interno no servidor", 500)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

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
