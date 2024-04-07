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
		fmt.Println("1. Criar tarefa")
		fmt.Println("2. Listar tarefas")
		fmt.Println("3. Listar uma tarefas")
		fmt.Println("4. Excluir tarefa")
		fmt.Println("5. Sair")
		fmt.Print("Opção: ")

		opcao := lerOpcao()

		switch opcao {
		case 1:
			criarTarefa(&listaDeTarefas)
		case 2:
			listarTarefas(listaDeTarefas)
		case 3:
			listaUmaTarefa(listaDeTarefas)
		case 4:
			removeTarefa(&listaDeTarefas)
		case 5:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}

func lerOpcao() int {
	leitor := bufio.NewReader(os.Stdin)
	opcaoStr, _ := leitor.ReadString('\n')
	opcaoStr = strings.TrimSpace(opcaoStr)
	opcao, _ := strconv.Atoi(opcaoStr)
	return opcao
}

// Create
func criarTarefa(listaDeTarefas *[]Tarefa) {
	fmt.Println("Criando uma nova tarefa:")
	leitor := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o nome da tarefa: ")
	nome, _ := leitor.ReadString('\n')
	nome = strings.TrimSpace(nome)

	fmt.Print("Digite a descrição da tarefa: ")
	descricao, _ := leitor.ReadString('\n')
	descricao = strings.TrimSpace(descricao)

	novaTarefa := Tarefa{
		id:            len(*listaDeTarefas) + 1,
		nome:          nome,
		descricao:     descricao,
		dataDeCriacao: time.Now(),
	}

	*listaDeTarefas = append(*listaDeTarefas, novaTarefa)

	fmt.Println("Tarefa criada com sucesso!")
}

// Read All
func listarTarefas(listaDeTarefas []Tarefa) {
	if len(listaDeTarefas) == 0 {
		fmt.Println("Não há tarefas cadastradas.")
		return
	}

	fmt.Println("Listando tarefas:")
	for _, tarefa := range listaDeTarefas {
		fmt.Printf("ID: %d\n", tarefa.id)
		fmt.Printf("Nome: %s\n", tarefa.nome)
		fmt.Printf("Descrição: %s\n", tarefa.descricao)
		fmt.Printf("Data de criação: %s\n", tarefa.dataDeCriacao.Format("02/01/2006"))
		fmt.Println("----------------------------------------")
	}
}

// Read
func listaUmaTarefa(listaDeTarefas []Tarefa) *Tarefa {
	fmt.Print("Digite o ID da tarefa a ser listada: ")
	leitor := bufio.NewReader(os.Stdin)
	idStr, _ := leitor.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	for _, tarefa := range listaDeTarefas {
		if tarefa.id == id {
			fmt.Println(tarefa.id)
			fmt.Println(tarefa.nome)
			fmt.Println(tarefa.descricao)
			fmt.Println(tarefa.dataDeCriacao)
		}
	}

	return nil
}

// Delete
func removeTarefa(listaDeTarefas *[]Tarefa) {
	if len(*listaDeTarefas) == 0 {
		fmt.Println("Não há tarefas cadastradas para serem removidas.")
		return
	}

	fmt.Print("Digite o ID da tarefa a ser removida: ")
	leitor := bufio.NewReader(os.Stdin)
	idStr, _ := leitor.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID inválido. Por favor, digite um número válido.")
		return
	}

	for i, tarefa := range *listaDeTarefas {
		if tarefa.id == id {
			*listaDeTarefas = append((*listaDeTarefas)[:i], (*listaDeTarefas)[i+1:]...)
			fmt.Println("Tarefa removida com sucesso!")
			return
		}
	}

	fmt.Println("Nenhuma tarefa encontrada com o ID especificado.")
}
