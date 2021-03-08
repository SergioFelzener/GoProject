package main

// como podemos usar heranca
import "fmt"

type people struct {
	name  string
	age   uint8
	email string
}

type student struct {
	people
	curse      string
	university string
}

func main() {
	fmt.Println("Herança só que não")

	p1 := people{name: "Maria", age: 40, email: "maria@email.com"}
	p2 := people{"Jose", 32, "jose@email.com"}

	var st1 student
	st1.name = "Teste"
	st1.age = 33
	st1.email = "joao@teste.com"
	st1.curse = "Informática"
	st1.university = "Senac"

	st2 := student{p2, "Matemática", "Senac universidade"}

	fmt.Println(st1)
	fmt.Println("---------")
	fmt.Println(st2)
	fmt.Println(st2.name)
	fmt.Println(p1)
}
