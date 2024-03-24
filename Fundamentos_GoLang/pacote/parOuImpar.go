package main

import "fmt"

func parOuImpar(numeros []int) {
	for _, valor := range numeros {
		if valor%2 == 0 {
			fmt.Printf("O número %d é par\n", valor)
		} else {
			fmt.Printf("O número %d é ímpar\n", valor)
		}
	}
}
