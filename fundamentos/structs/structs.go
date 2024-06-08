package main

import "fmt"

type endereco struct {
	numero     uint8
	logadrouro string
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Structs")

	var u usuario
	u.idade = 19
	u.nome = "Rafael"
	fmt.Println(u)

	enderecoEx := endereco{100, "Av. Brasil"}
	usuario2 := usuario{"Rafael", 19, enderecoEx}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 19}
	fmt.Println(usuario3)
}
