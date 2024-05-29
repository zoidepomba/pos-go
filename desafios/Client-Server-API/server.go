package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type StructCotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
	EURBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"EURBRL"`
	BTCBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"BTCBRL"`
}

func main() {
	http.HandleFunc("/cotacao", handleCotacao)
	http.ListenAndServe(":8080", nil)
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)	
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, "Error ao  criar a requisição", http.StatusInternalServerError)
		log.Println("Erro  ao  criar a requisição", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Erro  ao fazer  a requisição externa",  http.StatusInternalServerError)
		log.Println("erro ao   fazer a requisição externa", err)
		return
	}
	defer resp.Body.Close()

	var s StructCotacao

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		http.Error(w, "Erro ao ler a resposta externa", http.StatusInternalServerError)
		log.Println("Erro ao ler a resposta externa", err)
		return
	}

	if err := json.Unmarshal(body, &s); err != nil {
		http.Error(w, "Erro ao decodificar a resposta do json", http.StatusInternalServerError)
		log.Println("Erro ao decodificar resposta JSON:", err)
		return 
	}

	select {
	case <-ctx.Done():
		log.Println("Request cancelada pelo client")
        http.Error(w, "Request cancelada pelo client", http.StatusRequestTimeout)
		return
	default:
	}

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(s); err != nil {
        http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
        log.Println("Erro ao codificar resposta JSON:", err)
        return
    }

    log.Println("Request processada com sucesso")
}


