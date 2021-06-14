package main

import (
	"log"
	"net/http"
)

// Define uma função para servir como handler. Será o handler inicial.
// Através do ResponseWriter, vamos usar a função Write que vai receber
// um slice of bytes com o texto "Snippetbox" que será enviado no corpo da resposta.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippetbox"))
}

func main() {

	// http.NewServeMux() inicializa um novo servemux
	mux := http.NewServeMux()
	// registra a função home como um handler da url "/"
	mux.HandleFunc("/", home)

	//Faz um log na tela, mostrando que o servidor está iniciando
	log.Println("Iniciando servidor na porta 4000")

	// A função http.ListenAndServe() inicia um servidor web.
	// Passamos dois parâmetros:
	// - O endereço TCP da rede que vamos escutar (no exemplo, porta :4000)
	// - O servemux que acabamos de criar.
	// Se http.ListenAndServe() retornar um erro, usamos log.Fatal() para
	//fazer um log da mensagem de erro e então sair.
	// Qualquer erro retornado por http.ListenAndServe é sempre non-nil
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

