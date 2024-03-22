package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 50, 60, 70}

	// 1. Imprime o comprimento (len) e a capacidade (cap) do slice, juntamente com seus valores
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// 2. Cria um slice vazio a partir de s[:0] e imprime seu comprimento (len) e capacidade (cap), juntamente com seus valores
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
}
