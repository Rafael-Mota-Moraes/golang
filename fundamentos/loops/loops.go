package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		// time.Sleep(time.Second / 2)
		i++
		fmt.Println(i)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	nomes := []string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Rafael",
		"sobrenome": "Moreira",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Loop infinito
	/*
		for {
			fmt.Println("INFINITO")
		}
	*/
}
