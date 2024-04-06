package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tarefa struct {
	id            int
	nome          string
	descricao     string
	dataDeCriacao time.Time
}

func main() {

	listaDeTarefas := []Tarefa{}

	for {
		fmt.Println("Selecione uma opção:")
		println("------------------------")
		fmt.Println("1. Criar tarefa")
		println("------------------------")
		fmt.Println("2. Listar tarefas")
		println("------------------------")
		fmt.Println("3. Sair")

		// Ler a entrada do usuário
		opcao := lerOpcao()

		// Executar ação com base na opção selecionada
		switch opcao {
		case 1:
			criarTarefa(&listaDeTarefas)

		case 2:
			listarTarefas(listaDeTarefas)
		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}

}

func lerOpcao() int {
	leitor := bufio.NewReader(os.Stdin)    // criando um leitor
	opcaoStr, _ := leitor.ReadString('\n') // quando o leitor achar uma quebra de linha no caso o enter ele vai para de capturar o que o usuário digitou
	opcaoStr = strings.TrimSpace(opcaoStr) // remove os espaçoes desnecessários no inico e no final
	opcao, _ := strconv.Atoi(opcaoStr)     // convertendo uma string para um inteiro
	return opcao                           // retornando a opção
}

func criarTarefa(listaDeTarefas *[]Tarefa) {
	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Digite o nome da tarefa:")
	nome, _ := leitor.ReadString('\n')
	println()

	fmt.Println("Digite a descrição da tarefa:")
	descricao, _ := leitor.ReadString('\n')
	println()

	novaTarefa := Tarefa{
		id:            len(*listaDeTarefas) + 1,
		nome:          nome,
		descricao:     descricao,
		dataDeCriacao: time.Now(),
	}

	*listaDeTarefas = append(*listaDeTarefas, novaTarefa)

	println("Tarefa Criada com sucesso!")

}

func listarTarefas(listaDeTarefas []Tarefa) {
	for _, tarefa := range listaDeTarefas {
		fmt.Println("Id: ", tarefa.id)
		fmt.Println("Nome: ", tarefa.nome)
		fmt.Println("Descrição: ", tarefa.descricao)
		fmt.Println("Data de criação: ", tarefa.dataDeCriacao.Format("02/01/2006"))
		fmt.Println("----------------------------------------")
	}
}
