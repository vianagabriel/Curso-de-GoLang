package main

import "fmt"

func main() {
	// Define uma função que retorna uma função
	greeting := func() {
		message := "Olá" // Variável capturada pela closure
		// Define uma função interna
		printGreeting := func() {
			fmt.Println(message)
		}
		// Chama a função interna
		printGreeting()
	}

	// Chama a função que imprime a saudação
	greeting()
}

//Uma closure em Go é uma função que pode referenciar variáveis definidas fora de seu próprio escopo.Em outras palavras,
//ela "fecha" ou captura o ambiente em que foi definida. Isso significa que uma closure pode acessar e manipular variáveis
//​​que estão fora de seu próprio corpo de função, mesmo após a função ter retornado.
