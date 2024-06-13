package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		Http é um protocolo de comunicação - base da comunicação web

		Cliente - Servidor

		Rotas
			URI - Identificador do recurso
			Método - GET, POST, PUT, DELETE
	*/
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá, tudo bom?"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user 1\nuser 2"))
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
