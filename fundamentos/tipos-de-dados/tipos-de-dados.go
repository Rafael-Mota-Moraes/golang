package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos numéricos inteiros
	// int8, int16, int32, int64
	var numero int8 = 100
	fmt.Println(numero)
	numero2 := 100 // int64 (depende da arquitetura)
	fmt.Println(numero2)

	var numero3 uint32 = 1 // Unsigned int (números sem sinal)
	fmt.Println(numero3)

	// alias
	var numero4 rune = 1234 // igual ao int32
	fmt.Println(numero4)

	// alias
	var numero5 byte = 1 // igual ao int8
	fmt.Println(numero5)

	// Tipos numericos reais
	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// Cadeia de caracteres
	var str string = "Meu nome é Rafael"
	fmt.Println(str)

	char := 'b' // caracter converte para tabela ASCII
	fmt.Println(char)

	// Valor zero
	var texto string
	fmt.Println(texto) // string vazia

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}
