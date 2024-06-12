// TESTE DE UNIDADE
package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Estrada ABC", "Estrada"},
		{"Rodovia ABC", "Rodovia"},
		{"Praça ABC", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s \n",
				tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}
