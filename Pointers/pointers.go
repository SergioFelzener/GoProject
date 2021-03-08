package main

import "fmt"

func main() {
	fmt.Println("Ponteiros\n")
	fmt.Println("------")

	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	fmt.Println("------")

	variable1++

	fmt.Println(variable1, variable2)

	// PONTEIRO = REFERENCIA DE MEMORIA
	// SALVA O ENDERECO DE MEMORIA NO QUAL ESTA A VARIAVEL EM QUESTAO
	fmt.Println("------")

	var variable3 int = 100

	var pointer *int

	fmt.Println(variable3, pointer)

	fmt.Println("------")

	variable3 = 1000
	pointer = &variable3 // usa & para atribuir o ponteiro que é o valor salvo no endereço de memoria
	fmt.Println(variable3, pointer)
	//lendo dentro do endereco de memoria usa o *
	fmt.Println(variable3, *pointer)

}
