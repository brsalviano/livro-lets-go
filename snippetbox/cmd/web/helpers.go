package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// O helper serverError escreve uma mensagem de erro e a stack trace no errorLog
// e então manda uma resposta genérica 500 Internal Server Error para o usuário
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//app.errorLog.Println(trace)
	app.errorLog.Output(2, trace) //Vai nos mostrar a linha que chamou o erro!

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// O helper clientError manda um status específico e a descrição correspondente para o usuário.
// Vamos usar isso futuramente para enviar respostas. Como 400 "Bad Request" quando tiver um problema
// na requisição que o usuário enviou
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Para ser consistente, nós também implementamos o notFound helper.
// Isto é um simples embrulho em volta do clientError que envia uma resposta 404 Not Found
// para o usuário
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
