package main

import "fmt"

func main() {
	fmt.Println("Maps")
	fmt.Println("----------")

	user := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "silva",
	}

	fmt.Println(user)
	fmt.Println(user["nome"])
	fmt.Println("----------")

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joao",
			"ultimo":   "Nadson",
		},
		"curso": {
			"nome":         "matematica",
			"universidade": "senac",
		},
	}

	fmt.Println(user2)
	fmt.Println(user2["nome"]["primeiro"])

	fmt.Println("----------")
	// apagando a chave
	delete(user2, "nome")
	fmt.Println(user2)
	fmt.Println("----------")

	//recolocando a chave deve atender a mesma estrutura

	user2["nome"] = map[string]string{
		"primeiro": "Anderson",
		"ultimo":   "Silva",
	}

	fmt.Println(user2)
}
