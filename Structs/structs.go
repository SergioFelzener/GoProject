package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u user
	u.name = "Sergio"
	u.age = 22

	fmt.Println(u)

	u2 := user{"Andre", 26}
	fmt.Println(u2)

	//passando apenas um dos campos
	u3 := user{name: "Jose"}

	fmt.Println(u3)

	u4 := user{age: 12}
	fmt.Println(u4)

}
