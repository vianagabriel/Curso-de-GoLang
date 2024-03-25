package exercicio

import "fmt"

func BuscaBinaria(numeros []int, alvo int) {

	inicio := 0
	final := len(numeros) - 1
	numeroAlvo := alvo
	tentativas := 0

	for inicio <= final {
		meio := (inicio + final) / 2

		if numeroAlvo == numeros[meio] {
			fmt.Println("Encontrado", numeroAlvo)
			tentativas++
			break
		} else if numeros[meio] < numeroAlvo {
			inicio = meio + 1
			tentativas++
		} else {
			final = meio - 1
			tentativas++
		}
	}

	fmt.Printf("Número %d encontrado em %d tentativas\n", numeroAlvo, tentativas)
}
