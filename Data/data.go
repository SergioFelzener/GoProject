package main

import (
	"errors"
	"fmt"
)

func main() {
	number := 10000000
	fmt.Println(number)

	var number2 uint32 = 100000 // unsygned int sem sinal
	fmt.Println(number2)

	//alias
	// INT32 = RUNE
	var number3 rune = 12345
	fmt.Println(number3)

	//BYTE = UINT8
	var number4 byte
	fmt.Println(number4)

	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	var realNumber2 float64 = 1230000000.45
	fmt.Println(realNumber2)

	//tipo float so quando usa inferencia base arquitetura do sistema
	realNumber3 := 1298000.00
	fmt.Println(realNumber3)

	// FIM DOS NUMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// buscando valor numérico na tabela ASCII usa aspas simples em 1 caracter
	char := 'B'
	fmt.Println(char)

	// FIM STRING

	// booleano
	var boolean bool = true
	fmt.Println(boolean)

	var boolean2 bool = false
	fmt.Println(boolean2)

	// quando nao declara valor padrao é false
	var boolean3 bool
	fmt.Println(boolean3)

	var error error = errors.New("Erro Interno")
	fmt.Println(error)

}
