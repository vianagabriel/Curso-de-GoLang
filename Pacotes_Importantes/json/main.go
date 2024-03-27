package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Definindo a estrutura (struct) Pessoa para representar informações sobre uma pessoa
type Pessoa struct {
	Nome      string // Nome da pessoa
	Sobrenome string // Sobrenome da pessoa
	Idade     int    // Idade da pessoa
	Profissao string // Profissão da pessoa
}

func main() {

	// Criando uma instância da estrutura Pessoa
	cliente := Pessoa{Nome: "Gabriel", Sobrenome: "Viana", Idade: 25, Profissao: "Desenvolvedor"}

	// Convertendo a instância Pessoa para formato JSON
	res, err := json.Marshal(cliente) // Marshal converte a estrutura em []byte
	if err != nil {
		panic(err) // Se houver erro, exibe e encerra o programa
	}

	// Imprimindo o JSON resultante
	fmt.Println("Json ===>", string(res))

	// Criando um novo codificador JSON que escreve no os.Stdout (saída padrão)
	enconder := json.NewEncoder(os.Stdout).Encode(cliente)
	if enconder != nil {
		panic(enconder) // Se houver erro, exibe e encerra o programa
	}

	// ===================================================================

	// Definindo um formato de JSON como []byte
	formatoJson := []byte(`{"Nome": "Rebeca", "Sobrenome": "Viana", "Idade": 21, "Profissao": "DevOps"}`)
	var cliente2 Pessoa // Criando uma nova instância da estrutura Pessoa

	// Convertendo o JSON para a estrutura Pessoa
	err = json.Unmarshal(formatoJson, &cliente2) // Unmarshal converte []byte em estrutura
	if err != nil {
		fmt.Println(err) // Se houver erro, exibe o erro
	}

	// Imprimindo o nome da pessoa
	fmt.Println(cliente2.Nome)

}
