package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

type StructCotacao struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
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
	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//_, err = file.WriteString(fmt.Sprintf("Dolar: %s"))
	io.Copy(os.Stdout, res.Body)
	var cotacao StructCotacao
	println(cotacao.USDBRL.Bid)
	//_, err = file.WriteString(fmt.Sprintf("dolar", res.Body))
}
