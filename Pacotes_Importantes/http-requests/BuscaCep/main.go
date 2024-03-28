package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct { // CRIANDO UMA STRUCT E SETANDO TAGS PARA ELA
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
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/") // BUSCANDO AS INFORMAÇOES DA API DO VIACEP ATRAVES DA LIB NET/HTTP COM A FUNCAO GET
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)

		}
		defer req.Body.Close() // FECHANDO A REQUISICAO

		res, err := io.ReadAll(req.Body) // UTILIZANDO A LIB IO E A FUNCAO ReadAll PARA FAZER A LEITURA DA RESPOSTA
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)

		}

		var data ViaCEP                  // CRIANDO UMA STRUCT DO TIPO VIA CEP PARA PODER SALVAR OS DADOS RECEBIDOS PELO BODY
		err = json.Unmarshal(res, &data) // CONVERTENDO A RESPOSTA PARA UMA STRUCT E SALVANDO ELA NO LOCAL EM MEMORIA ONDE ESTÁ A VARIAVEL DATA
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)

		}

		file, err := os.Create("arquivo.txt") // CRIANDO UM ARQUIVO DE TEXTO COM A LIB 'os' USANDO O METODO CREATE
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)

		}

		defer file.Close() // FECHANDO O ARQUIVO FILE
		_, err = file.WriteString(fmt.Sprintf("CEP: %s,\nLocalidade: %s,\nUF: %s,\nLogradouro: %s\n", data.Cep, data.Localidade, data.Uf, data.Logradouro))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
	}
}
