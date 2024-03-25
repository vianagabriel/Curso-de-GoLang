package main

import (
	"fmt"
)

func main() {
	a := 10                // declarando a variável
	var ponteiro *int = &a // Declaração do ponteiro, ele está apontando para o endereço de memória de 'a'

	fmt.Println("Endereço de memória --->", ponteiro)

	fmt.Println("Valor armazenado no endereço --->", *ponteiro)

	*ponteiro = 20 // indo no endereço apontado pelo ponteiro que no caso é o endereço de 'a' e modificando o valor dela de 10 para 20

	fmt.Println("Valor alterado:", a)

	b := &a // declarando a variável 'b' e ela está apontando para o endereço de memória de 'a'

	fmt.Println(*b) // buscando o valor que está salvo no endereço de memória apontado por 'b'

	*b = 30 // indo no endereço de memória apontado por 'b' e mudando o valor para 30
	// OBS: 'b' está apontando para 'a' então quando se modifica o valor de 'b' também modifica o de 'a'

}
