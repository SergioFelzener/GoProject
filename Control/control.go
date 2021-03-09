package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")
	fmt.Println("------------------")

	number := -1

	if number > 15 {
		fmt.Println("Número maior que 15")
	} else {
		fmt.Println("Número menor ou igual a 15")
	}
	n2 := 0
	// criando variavel no if init ela so vale dentro do escopo do if
	fmt.Println("------------------")
	if n2 := number; n2 > 0 {
		fmt.Println("Número maior que Zero")
		fmt.Println(n2)
	} else {
		fmt.Println("Número menor que Zero")
		fmt.Println(n2)
	}

	fmt.Println("------------------")
	fmt.Println(n2)

}
