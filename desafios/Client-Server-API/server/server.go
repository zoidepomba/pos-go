package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type StructCotacao struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {

	criaTabela()

	http.HandleFunc("/cotacao", handleCotacao)
	log.Printf("Servidor rodando na porta :8080\n")
	log.Printf("Servidor rodando no path /cotacao ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func criaTabela() {
	log.Printf("Criando a tabela no banco de dados")
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados: %v", err)
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cotacoes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			valor TEXT,
			data TEXT
		)
	`)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela: %v", err)
	}
}

func execRequest(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cotacao StructCotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return "", err
	}

	return cotacao.USDBRL.Bid, nil
}

func saveCotacaoDB(ctx context.Context, db *sql.DB, valor string) error {
	query := "INSERT INTO cotacoes (valor, data) VALUES (?, datetime('now'))"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, valor)
	return err
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")

	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	defer log.Println("Request finalizada")

	req, err := execRequest(ctx)
	if err != nil {
		http.Error(w, "Error ao  criar a requisição", http.StatusInternalServerError)
		log.Println("Erro  ao  criar a requisição", err)
		return
	}

	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		log.Printf("Erro ao abrir a conexão com o banco de dados: %v", err)
		http.Error(w, "Erro ao abrir o banco de dados", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	if err := saveCotacaoDB(ctxDB, db, req); err != nil {
		http.Error(w, "Erro ao salvar cotação", http.StatusInternalServerError)
		log.Println("Erro ao salvar cotação: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": req})
	log.Println("Request processada com sucesso")
}
