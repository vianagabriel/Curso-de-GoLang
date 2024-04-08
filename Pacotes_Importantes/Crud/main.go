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
		fmt.Println("3. Listar uma tarefa")
		fmt.Println("4. Excluir tarefa")
		fmt.Println("5. Editar tarefa")
		fmt.Println("6. Sair")
		println()
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
			editaTarefa(&listaDeTarefas)
		case 6:
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

func formataApresentecao(tarefa *Tarefa) {
	fmt.Printf("ID: %d\n", tarefa.id)
	fmt.Printf("Nome: %s\n", tarefa.nome)
	fmt.Printf("Descrição: %s\n", tarefa.descricao)
	fmt.Printf("Data de criação: %s\n", tarefa.dataDeCriacao.Format("02/01/2006"))
	fmt.Println("----------------------------------------")
	println()
}

// ======================================================C.R.U.D.==================================================================

func criarTarefa(listaDeTarefas *[]Tarefa) {
	println()

	fmt.Println("Criando uma nova tarefa:")
	println("============================")
	leitor := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o nome da tarefa: ")
	nome, _ := leitor.ReadString('\n')
	nome = strings.TrimSpace(nome)

	println("===========================")

	fmt.Print("Digite a descrição da tarefa: ")
	descricao, _ := leitor.ReadString('\n')
	descricao = strings.TrimSpace(descricao)

	println("===========================")

	novaTarefa := Tarefa{
		id:            len(*listaDeTarefas) + 1,
		nome:          nome,
		descricao:     descricao,
		dataDeCriacao: time.Now(),
	}

	*listaDeTarefas = append(*listaDeTarefas, novaTarefa)

	fmt.Println("Tarefa criada com sucesso!")
	println()
}

func listarTarefas(listaDeTarefas []Tarefa) {
	if len(listaDeTarefas) == 0 {
		fmt.Println("Não há tarefas cadastradas.")
		return
	}

	println()
	fmt.Println("Listando tarefas:")
	println()

	for _, tarefa := range listaDeTarefas {
		formataApresentecao(&tarefa)
	}
}

func listaUmaTarefa(listaDeTarefas []Tarefa) *Tarefa {
	println()

	fmt.Print("Digite o ID da tarefa a ser listada: ")

	leitor := bufio.NewReader(os.Stdin)
	idStr, _ := leitor.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	println()

	for _, tarefa := range listaDeTarefas {
		if tarefa.id == id {
			formataApresentecao(&tarefa)
			return &tarefa
		}
	}

	fmt.Println("Nenhuma tarefa encontrada com o ID especificado.")
	return nil
}

func editaTarefa(listaDeTarefas *[]Tarefa) {
	println()
	fmt.Print("Digite o ID da tarefa a ser editada: ")
	leitor := bufio.NewReader(os.Stdin)
	idStr, _ := leitor.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)
	println()

	for {
		fmt.Println("Selecione uma opção:")
		fmt.Println("1. Editar o nome")
		fmt.Println("2. Editar a descrição")
		fmt.Println("3. Editar nome e descrição")
		fmt.Println("4. Sair")
		fmt.Print("Opção: ")

		opcao := lerOpcao()
		switch opcao {
		case 1, 2, 3:
			var nome, descricao string
			if opcao == 1 || opcao == 3 {
				println()
				leitor := bufio.NewReader(os.Stdin)
				fmt.Print("Digite o nome da tarefa: ")
				nome, _ = leitor.ReadString('\n')
				nome = strings.TrimSpace(nome)
			}
			if opcao == 2 || opcao == 3 {
				println()
				leitor := bufio.NewReader(os.Stdin)
				fmt.Print("Digite a descrição da tarefa: ")
				descricao, _ = leitor.ReadString('\n')
				descricao = strings.TrimSpace(descricao)
			}

			for i, tarefa := range *listaDeTarefas {
				if tarefa.id == id {
					if opcao == 1 || opcao == 3 {
						(*listaDeTarefas)[i].nome = nome
					}
					if opcao == 2 || opcao == 3 {
						(*listaDeTarefas)[i].descricao = descricao
					}
					fmt.Println("Tarefa editada com sucesso!")
					return
				}
			}
		case 4:
			println()
			fmt.Println("Saindo...")
			return
		default:
			println()
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}

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
