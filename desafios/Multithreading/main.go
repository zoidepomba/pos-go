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
	fmt.Println("Buscando CEP")
	cep := "38412392"
	buscarEndereco(cep)
}

func buscarEndereco(cep string) {
	client := &http.Client{Timeout: time.Second}
	chBrasilAPI := make(chan *StructBrasilApi)
	chViaCEP := make(chan *StructViaCEP)
	start := time.Now()

	go BrasilApi(client, cep, chBrasilAPI, start)
	go ViaCEP(client, cep, chViaCEP, start)

	select {
	case result := <-chBrasilAPI:
		if result != nil{
			fmt.Println("BrasilAPI Respondeu primeiro !")
			fmt.Printf("Cep: %s, Estado: %s, Cidade: %s, Bairro: %s, Rua: %s\n", result.Cep, result.State, result.City, result.Neighborhood, result.Street)
		} else {
			fmt.Println("BrasilAPi retornou resultado invalido")
		}
	case result := <-chViaCEP:
		if result != nil {
			fmt.Println("ViaCEP respondeu primeiro !")
			fmt.Printf("Cep: %s, Logradouro: %s, Bairro: %s, Localidade: %s, UF: %s\n", result.Cep, result.Logradouro, result.Bairro, result.Localidade, result.Uf)
		} else{
			fmt.Println("ViaCEP retornou um resultado invalido !")
		}
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: TIMEOUT")
	}

}

func ViaCEP(client *http.Client, cep string, ch chan<- *StructViaCEP, start time.Time) {
	resp, err := client.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Println("Erro na API ViaCEP:", err)
		ch <- nil
		return
	}
	defer resp.Body.Close() //fechando a request

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro na leitura da resposta ViaCEP:", err)
		ch <- nil
		return
	}
	var s StructViaCEP

	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Erro no Unmarshal ViaCEP:", err)
		ch <- nil
		return
	}
	duration := time.Since(start)
	fmt.Printf("Tempo para resposta ViaCEP: %s\n", duration)
}

func BrasilApi(client *http.Client, cep string, ch chan<- *StructBrasilApi, start time.Time) {
	resp, err := client.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		fmt.Println("Erro na API BrasilAPI:", err)
		ch <- nil
		return
	}
	defer resp.Body.Close() //fechando a request

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro na leitura da resposta BrasilAPI:", err)
		ch <- nil
		return
	}
	var s StructBrasilApi

	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Erro no Unmarshal BrasilAPI:", err)
		ch <- nil
		return
	}
	duration := time.Since(start)
	fmt.Printf("Tempo para resposta BrasilApi: %s\n", duration)

	ch <- &s
}
