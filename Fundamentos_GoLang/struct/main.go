package main

import (
	"fmt"
)

type Prestador struct {
	Nome      string
	User_Name string
	Ativo     bool
}

func main() {
	funcionario := Prestador{
		Nome:      "Gabriel",
		User_Name: "gabriel.viana",
		Ativo:     true,
	}

	fmt.Println(funcionario)

}
