package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := 10; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
		fmt.Println(outroNumero)
	} else if numero < -10 {
		fmt.Println("Número é maior que -10")
	} else {
		fmt.Println("Está entre zero e -10")

	}
	// fmt.Println(outroNumero)
}
