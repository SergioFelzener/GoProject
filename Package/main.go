package main

import (
	"fmt"
	"modulo/aux"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Teste de Impressao")
	aux.Escrever()

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)

}
