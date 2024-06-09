package main

import "fmt"

var n int

// Executa antes da main
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println(n)
}
