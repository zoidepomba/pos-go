package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type StructBrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type StructViaCEP struct {
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
	fmt.Println("iniciando o programa")
	cep := "38412392"
	chamada, err := ViaCEP(cep)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Dados da API Brasil:\nCep:%s ", chamada.Cep )
	fmt.Printf(chamada.Cep)
}

func ViaCEP(cep string) (*StructViaCEP, error) {
	c := http.Client{Timeout: time.Second}
	resp, err := c.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //fechando a request

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var s StructViaCEP

	err = json.Unmarshal(body, &s)
	if err != nil{
		panic(err)
	}
	return &s, nil
}

func BrasilApi(cep string) (*StructBrasilApi, error) {
	c := http.Client{Timeout: time.Second}
	resp, err := c.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //fechando a request

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var s StructBrasilApi

	err = json.Unmarshal(body, &s)
	if err != nil{
		panic(err)
	}
	return &s, nil
}
