package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Faz uma requisição GET para a URL especificada e retorna uma resposta e um possível erro.
	req, err := http.Get("https://www.google.com")
	// Verifica se houve um erro na requisição.
	if err != nil {
		// Se houve um erro, o programa entra em estado de pânico e exibe o erro.
		panic(err)
	}
	// Adia o fechamento do corpo da resposta até o final desta função ser alcançado.
	defer req.Body.Close()

	// Lê o corpo da resposta e armazena os dados lidos em 'res'.
	res, err := io.ReadAll(req.Body)
	// Verifica se houve um erro na leitura do corpo da resposta.
	if err != nil {
		// Se houve um erro, o programa entra em estado de pânico e exibe o erro.
		panic(err)
	}

	// Converte os dados lidos em 'res' para uma string e os imprime.
	fmt.Println(string(res))
}
