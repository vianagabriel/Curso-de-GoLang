package main

import (
	"fmt"
	"os"
)

func main() {
	// Criando um arquivo de texto chamado "arquivo.txt"
	f, err := os.Create("arquivo.txt")
	if err != nil { // Verifica se ocorreu algum erro ao criar o arquivo
		panic(err) // Em caso de erro, interrompe o programa e exibe o erro
	}

	// Gravando a string "Gabriel Viana" no arquivo
	tamanho, err := f.Write([]byte("Gabriel Viana"))
	if err != nil { // Verifica se ocorreu algum erro ao gravar no arquivo
		panic(err) // Em caso de erro, interrompe o programa e exibe o erro
	}

	// Imprimindo o tamanho do arquivo criado
	fmt.Printf("Arquivo criado com sucesso! Tamanho de %d bytes\n", tamanho)

	// Fechando o arquivo após concluir as operações
	f.Close()

	//------------------------------ Leitura

	// Lendo o conteúdo do arquivo "arquivo.txt" e armazenando-o na variável 'arquivo'
	arquivo, err := os.ReadFile("arquivo.txt")

	// Verificando se ocorreu algum erro durante a leitura do arquivo
	if err != nil {
		// Se houver um erro, o programa entra em pânico e exibe o erro
		panic(err)
	}

	// Convertendo o conteúdo do arquivo de bytes para uma string e imprimindo-o no terminal
	fmt.Println(string(arquivo))
}
