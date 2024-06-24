package controllers

import "net/http"

// CarregarTelaDeLogin renderiza a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tela de login"))
}
