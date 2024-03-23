package main

import "fmt"

func Soma[T int | float64](m map[string]T) T {
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
