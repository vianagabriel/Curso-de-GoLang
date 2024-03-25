package main

import (
	"bufio"
	"fmt"
	"os"
)

// main é a função principal do programa.
func main() {
	// Criando um arquivo chamado "arquivo.txt" para escrita.
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Escrevendo um texto no arquivo.
	tamanho, err := f.Write([]byte("Go, uma linguagem de programação desenvolvida pela Google, ganhou destaque..."))
	if err != nil {
		panic(err)
	}

	// Informando ao usuário que o arquivo foi criado com sucesso e seu tamanho.
	fmt.Printf("Arquivo criado com sucesso! Seu tamanho é de %d bytes\n", tamanho)

	// Fechando o arquivo após terminar de escrever nele.
	f.Close()

	// --------------------------------- LEITURA ---------------------------------

	// Abrindo o arquivo "arquivo.txt" para leitura.
	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Inicializando um leitor de buffer para leitura do arquivo.
	reader := bufio.NewReader(arquivo)

	// Criando um buffer de 50 bytes para armazenar temporariamente os dados lidos do arquivo.
	buffer := make([]byte, 50)

	// Loop infinito para ler o arquivo pouco a pouco.
	for {
		// Lendo dados do arquivo e armazenando no buffer.
		n, err := reader.Read(buffer)
		if err != nil {
			break // Sai do loop se ocorrer um erro de leitura.
		}

		// Imprimindo os dados lidos do buffer como uma string.
		fmt.Println(string(buffer[:n]))
	}
}
