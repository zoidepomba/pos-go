/*package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCriaTabela(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS cotacoes").WillReturnResult(sqlmock.NewResult(1, 1))

	criaTabela()

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestExecRequest(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	randomBid := fmt.Sprintf("%.4f", 4.0+rand.Float64()*(6.0-4.0))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse := fmt.Sprintf(`{"USDBRL":{"bid":"%s"}}`, randomBid)
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	bid, err := execRequest(ctx, )
	assert.NoError(t, err)
	assert.Equal(t, randomBid, bid)
}

func TestSaveCotacaoDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	mock.ExpectPrepare("INSERT INTO cotacoes").
		ExpectExec().
		WithArgs("5.42").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = saveCotacaoDB(ctx, db, "5.42")
	assert.NoError(t, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestHandleCotacao(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse := `{"USDBRL":{"bid":"5.42"}}`
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	oldURL := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	defer func() { execRequestURL = oldURL }()
	execRequestURL = server.URL

	req := httptest.NewRequest("GET", "http://localhost:8080/cotacao", nil)
	w := httptest.NewRecorder()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("INSERT INTO cotacoes").
		WithArgs("5.42").
		WillReturnResult(sqlmock.NewResult(1, 1))

	handleCotacao(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "5.42", responseBody["bid"])

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}*/
