package main

import (
	"fmt"
)

func main() {
	radar(65, 60)

}

func radar(velocidade, limite int) string {
	if velocidade > limite {
		fmt.Printf("Multado por excesso de velocidade, velocidade = %d  limite da via %d", velocidade, limite)
	}

	return "Velocidade OK"
}
