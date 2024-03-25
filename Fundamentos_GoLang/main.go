package main

import "curso-go/Fundamentos_GoLang/exercicio"

func main() {
	numeros := []int{
		10, 15, 40, 60, 75, 85, 95, 110, 150, 220,
		225, 240, 330, 440, 555, 666, 765, 876, 998, 3000,
		3010, 3020, 3030, 3040, 3050, 3060, 3070, 3080, 3090, 3100,
	}
	exercicio.BuscaBinaria(numeros, 3050)
}
