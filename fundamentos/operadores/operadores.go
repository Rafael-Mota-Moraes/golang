package main

import "fmt"

func main() {
	// Aritméticos
	// + - / * %
	soma := 1 + 1
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 2
	modulo := 10 % 3
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Divisão:", divisao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Módulo:", modulo)

	/*
		Não funciona!! Em go, não se pode fazer nada com tipos diferentes...
		var numero1 int16 = 10
		var numero2 int32 = 25

		soma = numero1 + numero2
	*/

	// Atribuição
	var var1 string = "String" // Declaração normal
	var2 := "String 2"         // Usando inferência de tipos

	fmt.Println(var1, var2)

	// Relacionais
	// == >= > != <= <
	fmt.Println("1>2", 1 > 2)
	fmt.Println("1<2", 1 < 2)
	fmt.Println("1>=2", 1 >= 2)
	fmt.Println("1<=2", 1 <= 2)
	fmt.Println("1==2", 1 == 2)
	fmt.Println("1!=2", 1 != 2)

	// Lógicos
	// && || !
	fmt.Println("true && false", true && false)
	fmt.Println("true || false", true || false)
	fmt.Println("!true", !true)

	// Unários
	numero := 1
	numero++
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero -= 10
	fmt.Println(numero)
}
