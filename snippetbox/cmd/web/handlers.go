package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//A assinatura foi modificada para que seja um método de *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.partial.gohtml",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Como o home handler agora é um método de application
		// ele pode acessar seus campos, incluindo o logger de erros.
		// Vamos escrever as mensagens de log para o logger de
		// application em vez do logger padrão.
		app.errorLog.Println(err.Error())
		http.Error(w, "Erro interno no servidor", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		//Atualizamos o logger aqui também!
		app.errorLog.Println(err.Error())
		http.Error(w, "Erro interno no servidor", 500)
	}
}

//Mudamos a assinatura aqui também, para ser um método de *application.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Mostra um snippet específico com o ID %d...", id)
}

//Mudamos a assinatura aqui também, para ser um método de *application.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Método não permitido", 405)
		return
	}

	w.Write([]byte("Cria um novo snippet..."))
}
