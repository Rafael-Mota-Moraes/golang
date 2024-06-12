// TESTE DE UNIDADE
package enderecos

import (
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {
	enderecoParateste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParateste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado!\nEsperava: %s\nRecebeu: %s",
			tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}

}
