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
	funcionario.Desativar()
	funcionario.mudaNome("Rebeca")

}

func (prestador Prestador) Desativar() {
	prestador.Ativo = false

	fmt.Println(prestador)
}

func (prestador Prestador) mudaNome(nome string) {
	prestador.Nome = nome

	fmt.Println(prestador)
}
