package main

import (
	"fmt"
)

func main() {
	salarios := map[string]int{"Gabriel": 18000, "Rebeca": 18000}
	fmt.Println("Salário de Gabriel R$", salarios["Gabriel"])

	delete(salarios, "Gabriel")
	salarios["Keki"] = 50000

	fmt.Println("Salário de Keki R$", salarios["Keki"])

	salarios2 := map[string]int{}

	salarios2["Maria"] = 28000
	fmt.Println("Salário de Maria R$", salarios2["Maria"])

}
