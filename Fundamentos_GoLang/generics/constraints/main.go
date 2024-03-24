package main

import "fmt"

type Number interface { // posso criar uma interface para criar um novo tipo, nesse ex tenho uma interface Number
	// que recebe dois tipos: int | float64
	int | float64
}

func Soma[T Number](m map[string]T) T { //[T int | float64](m map[string]T) T  em vez de passar desse jeito passo o T com tipo Number
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func main() {
	numeros := map[string]int{"num1": 1, "num2": 3, "num3": 6}
	numeros2 := map[string]float64{"num1": 5.5, "num2": 3.3, "num3": 6.5}

	fmt.Println(Soma(numeros))
	fmt.Println(Soma(numeros2))
}
