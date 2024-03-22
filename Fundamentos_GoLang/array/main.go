package main

import "fmt"

func main() {
	// Declarando e inicializando um array de inteiros com 4 elementos
	arrNumeros := [4]int{10, 20, 30, 40}

	// Iterando sobre o array usando um loop for
	for i := 0; i < len(arrNumeros); i++ {
		// Imprimindo o índice atual e o valor correspondente no array
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, arrNumeros[i])
	}
}
