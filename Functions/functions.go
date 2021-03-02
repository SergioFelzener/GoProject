package main

import (
	"fmt"
)

func sum(n3 int8, n4 int8) int8 {
	return n3 + n4
}

// funcao com mais de um retorno
func calc(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	// usando a soma
	sum := sum(4, 6)
	fmt.Println(sum)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Texto da funçao 1")
	fmt.Println(result)

	// retorno funçao com 2 retornos
	resultSum, resultSub := calc(2, 2)
	fmt.Println(resultSum, resultSub)

	// ignorando um retorno com underline _
	resultSum2, _ := calc(2, 2)
	fmt.Println(resultSum2)

	// retornando apenas segundo return
	_, resultSub2 := calc(2, 2)
	fmt.Println(resultSub2)
}
