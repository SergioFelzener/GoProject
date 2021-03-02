package main

import "fmt"

func main() {
	// Aritimeticos
	// + , - , / , * , %

	sum := 2 + 3
	sub := 2 - 3
	div := 2 / 3
	mult := 2 * 3
	modulo := 5 % 2

	fmt.Println(sum, sub, div, mult, modulo)

	var num1 int16 = 10
	var num2 int32 = 25

	sum2 := num1 + int16(num2)
	fmt.Println(sum2)

	//FIM

	// Atribuição

	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"
	fmt.Println(variavel1, variavel2)
	fmt.Println("---------------------")

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println("---------------------")

	//operadores logicos

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println("---------------------\n") // pulando linha

	// operadores unários
	number := 10
	number++
	number += 15 // number = number + 15

	fmt.Println(number)

}
