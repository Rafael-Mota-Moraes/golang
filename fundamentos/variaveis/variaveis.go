package main

import "fmt"

func main() {
	// Strings
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		var3 string = "var3"
		var4 string = "var4"
	)

	fmt.Println(var3, var4)

	const constante1 string = "Constante"
	fmt.Println(constante1)
}
