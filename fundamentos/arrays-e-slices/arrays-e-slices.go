package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]int
	array1[0] = 10
	array1[1] = 20
	array1[2] = 30
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3"}
	array2[3] = "Posição 4"
	array2[4] = "Posição 5"
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15} // Parece com o array mas não é...
	fmt.Println(slice)
	slice = append(slice, 20)
	fmt.Println(slice)

	// TypeOf
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("---ARRAYS INTERNOS---")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	fmt.Println(slice3)

	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacidade
}
