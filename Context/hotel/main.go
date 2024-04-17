package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Criação de um contexto de fundo (Background)
	ctx := context.Background()

	// Define um prazo de 3 segundos para o contexto usando WithTimeout
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel() // Garante que a função cancel será chamada após a conclusão da função main

	// Chama a função para reservar o hotel, passando o contexto
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		// Se o contexto for cancelado antes do prazo, imprime mensagem de cancelamento
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(3 * time.Second):
		// Se a reserva for bem-sucedida dentro do prazo, imprime mensagem de sucesso
		fmt.Println("Hotel booked.")
	}
}
