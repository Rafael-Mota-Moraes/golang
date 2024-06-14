package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para struct"))
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
	}
	defer db.Close()

	// Prepare statement
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	// status codes
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso id: %d", idInserido)))
}
