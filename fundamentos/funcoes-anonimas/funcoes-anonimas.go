package main

import "fmt"

func main() {

	texto := func(texto string) string {
		return fmt.Sprintf("recebido -> %s", texto)
	}("Parâmetro")

	fmt.Println(texto)
}
