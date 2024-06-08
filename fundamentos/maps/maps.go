package main

import "fmt"

func main() {
	fmt.Println("MAPS")
	usuario := map[string]string{
		"nome":      "Rafael",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {"nome": "engenharia"},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
