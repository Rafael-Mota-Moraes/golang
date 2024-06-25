package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPI representa a resposta de erro da AP
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Cont-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// TratarStatusCodeErro trata as requisições com status code 400 ou superior
func TratarStatusCodeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
