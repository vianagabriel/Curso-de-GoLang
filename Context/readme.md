# Utilizando Contexto em Go (Golang)

O pacote `context` em Go é usado para transmitir sinais de cancelamento e prazo de execução entre goroutines. Isso é particularmente útil em operações concorrentes, como em servidores HTTP, onde é importante controlar o tempo de execução das solicitações.

## O que é um contexto?

Um contexto (`Context`) em Go é uma estrutura de dados que contém informações sobre o cancelamento de uma operação e seu prazo de execução. Ele é usado para transmitir sinais de cancelamento e prazo de execução entre partes de um programa.

## Por que usar contexto?

- **Cancelamento de operações**: Permite cancelar operações em andamento quando necessário, por exemplo, para evitar bloqueios em chamadas de rede.
- **Prazo de execução**: Define um limite de tempo para a execução de uma operação, útil para garantir que uma operação não demore indefinidamente para ser concluída.

## Como usar contexto em Go

1. **Criando um contexto**: O contexto principal (`context.Background()`) é usado como ponto de partida. Você pode criar um novo contexto com `context.WithCancel()` ou `context.WithTimeout()` para especificar um cancelamento ou prazo de execução.

2. **Transmitindo o contexto**: Passe o contexto como argumento para as funções que precisam acessá-lo. As funções que aceitam um contexto geralmente têm um parâmetro `context.Context`.

3. **Verificando o cancelamento**: Verifique o status de cancelamento do contexto usando `ctx.Done()`. Isso retorna um canal que é fechado quando o contexto é cancelado.

4. **Lidando com o prazo de execução**: Se você definiu um prazo de execução, o contexto será cancelado automaticamente quando o prazo expirar. Você pode usar `select` para aguardar a conclusão da operação ou o cancelamento do contexto.

## Exemplo de código

Aqui está um exemplo simples de como usar contexto em Go:

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Criando um contexto com prazo de execução de 2 segundos
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Garante que o contexto seja cancelado ao sair da função

    // Executando uma operação que pode demorar
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Operação concluída com sucesso!")
    case <-ctx.Done():
        fmt.Println("Operação cancelada devido ao prazo de execução.")
    }
}
