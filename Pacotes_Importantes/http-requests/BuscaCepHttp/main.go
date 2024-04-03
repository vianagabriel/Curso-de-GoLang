package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil) // Criando um servidor com http.ListenAndServe

}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) { // passando como parametro da funcao um w que é do tipo http.ResponseWriter que é a resposta e um r do tipo *http.Request que é a requisicao
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cep)

}

func BuscaCep(cep string) (*ViaCEP, error) {
	req, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}

	defer req.Body.Close()

	resBody, error := io.ReadAll(req.Body)
	if error != nil {
		return nil, error
	}

	var dataCep ViaCEP
	error = json.Unmarshal(resBody, &dataCep)
	if error != nil {
		return nil, error
	}

	return &dataCep, nil
}
