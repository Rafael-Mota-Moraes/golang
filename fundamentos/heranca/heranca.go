package main

import "fmt"

// O mais próximo que o go chega do conceito de herança

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "IFSUL"}
	fmt.Println(e1.nome)
	fmt.Println(e1.sobrenome)
	fmt.Println(e1)
}
