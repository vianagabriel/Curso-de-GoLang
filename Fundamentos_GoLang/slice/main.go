package main

import "fmt"

func main() {
	// Criando um slice com valores iniciais
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice de números:", numeros)

	// Acessando elementos de um slice
	fmt.Println("Primeiro elemento:", numeros[0])
	fmt.Println("Segundo elemento:", numeros[1])

	// Modificando elementos de um slice
	numeros[0] = 10
	fmt.Println("Slice após modificar o primeiro elemento:", numeros)

	// Adicionando um novo elemento ao slice
	numeros = append(numeros, 6)
	fmt.Println("Slice após adicionar um novo elemento:", numeros)

	// Criando um slice vazio
	var outroSlice []int
	fmt.Println("Slice vazio:", outroSlice)
}
