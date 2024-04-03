package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		// Realiza uma requisição HTTP para obter informações do CEP através da API do ViaCEP.
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
			return
		}
		defer req.Body.Close() // Garante que o corpo da requisição seja fechado após o uso.

		// Lê a resposta da requisição HTTP.
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
			return
		}

		// Converte os dados JSON recebidos em uma estrutura ViaCEP.
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
			return
		}

		// Cria um arquivo de texto para armazenar as informações do CEP.
		file, err := os.Create("arquivo.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
			return
		}
		defer file.Close() // Garante que o arquivo seja fechado após o uso.

		// Escreve as informações do CEP no arquivo.
		_, err = file.WriteString(fmt.Sprintf("CEP: %s,\nLocalidade: %s,\nUF: %s,\nLogradouro: %s\n", data.Cep, data.Localidade, data.Uf, data.Logradouro))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
			return
		}
	}
}
