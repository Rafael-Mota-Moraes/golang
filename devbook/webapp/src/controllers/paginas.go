package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/utils"
)

// CarregarTelaDeLogin renderiza a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario vai renderizar a página de cadastro de usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPagina principal carrega a home do site
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := http.Get(url)

	fmt.Println(response.StatusCode, erro)
	utils.ExecutarTemplate(w, "home.html", nil)
}
