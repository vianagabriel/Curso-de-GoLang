package main

import (
	"errors"
	"fmt"
)

func main() {
	velocidade, limite, err := radar(40, 50)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Limite da via: %d km\nVelocidade que você passou: %d km\n", limite, velocidade)

}

func radar(velocidade, limite int) (int, int, error) {
	if velocidade > limite {
		fmt.Printf("Limite da via: %d km\nVelocidade que você passou: %d km\n", limite, velocidade)
		return velocidade, limite, errors.New("Multado! Você excedeu o limite da via!")
	}

	return velocidade, limite, nil
}
