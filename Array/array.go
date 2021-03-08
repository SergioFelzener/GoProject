package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	//populando array

	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"position 1", "position 2", "position 3", "position 4", "position 5"}
	fmt.Println(array2)

	// mais flexivel tamanho na quantidade de intens que se passou para o array
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICES

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println("------")

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	fmt.Println("------")

	slice = append(slice, 18)
	fmt.Println(slice)

	fmt.Println("------")

	slice2 := array2[1:3]
	fmt.Println(slice2)

}
