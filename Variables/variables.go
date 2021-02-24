package main

import "fmt"

func main() {
	var var1 string = "variável 1"
	var2 := "Variável 2"
	fmt.Println(var1)
	fmt.Println(var2)

	var (
		var3 string = "Variável 3"
		var4 string = "Variavel 4"
	)

	fmt.Println(var3)
	fmt.Println(var4)

	var5, var6 := "Váriavel 5", "Variável 6"
	fmt.Println(var5, var6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}
