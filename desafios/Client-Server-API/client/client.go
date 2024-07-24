package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type StructCotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
	}
	var cotacao StructCotacao

	//println(cotacao.Bid)
	err = json.Unmarshal(resp, &cotacao)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v", err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dolar: %s", cotacao.Bid))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		return
	}
	fmt.Println("Arquivo criado com sucesso")
	fmt.Println("Cidade:", cotacao.Bid)

}
